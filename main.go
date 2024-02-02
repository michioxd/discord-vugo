package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

const version = "0.1"
const temporaryPath = "./.discord-vugo-tmp"

var fileNameTemp string = "dvugo_" + strconv.FormatInt(time.Now().UnixNano(), 10)
var ffmpegPath string = ""
var workingFilePath string = ""
var outputFilePath string = ""

func init() {
	fmt.Println("=======\ndiscord-vugo powered by michioxd\nversion: " + version + "\nhttps://github.com/michioxd/discord-vugo\n=======\n ")

	// init config
	// viper.SetConfigFile("./discord-vugo-config.yaml")

	viper.SetDefault("bot_token", "")
	viper.SetDefault("guild_channel_id", "")
	viper.SetDefault("file_max_second", "5")
	viper.SetDefault("use_proxy", false)
	viper.SetDefault("proxy_endpoint", "")

	viper.SetConfigName("discord-vugo-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if cfgErr := viper.ReadInConfig(); cfgErr != nil {
		if _, ok := cfgErr.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[WARN] Created configuration file, go to discord-vugo-config.yaml then edit something you need. You may check out the configuration at https://github.com/michioxd/discord-vugo/tree/main?tab=readme-ov-file#configuration")
			viper.WriteConfigAs("./discord-vugo-config.yaml")
			os.Exit(0)
		} else {
			panic(fmt.Errorf("Failed to read config file: %w", cfgErr))
		}
	}

	// check required cfg
	if viper.GetString("bot_token") == "" {
		panic(fmt.Errorf("Configuration error. `bot_token` is empty!"))
	} else if viper.GetString("guild_channel_id") == "" {
		panic(fmt.Errorf("Configuration error. `guild_channel_id` is empty!"))
	}

	if viper.GetBool("use_proxy") == true && viper.GetString("proxy_endpoint") == "" {
		panic(fmt.Errorf("Configuration error. `use_proxy` is enabled but `proxy_endpoint` is empty!"))
	}

	// check ffmpeg
	getFfmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		panic(fmt.Errorf("Cannot find FFmpeg (ffmpeg), make sure you have installed ffmpeg and add it to your `PATH`!"))
	} else {
		ffmpegPath = getFfmpegPath
	}

	// check input/output
	flag.StringVar(&workingFilePath, "i", "", "Input video path file. Example: -i \"D:/nggyu.mp4\"")
	flag.StringVar(&outputFilePath, "o", "./output.m3u8", "Output m3u8 path file. Default ./output.m3u8. Example: -o \"D:/nggyu-hls.m3u8\"")

	if workingFilePath == "" && viper.GetString("input_file") == "" {
		panic(fmt.Errorf("Please provide input file path argument or in configuration file!"))
	} else if workingFilePath == "" {
		workingFilePath = viper.GetString("input_file")
	}

	if outputFilePath == "" && viper.GetString("output_file") == "" {
		panic(fmt.Errorf("Output file path you have provided is empty! Please provide output file path argument or in configuration file!"))
	} else if outputFilePath == "" {
		workingFilePath = viper.GetString("output_file")
	}

}

func main() {
	fmt.Println("Creating session...")
	bot, botErr := discordgo.New("Bot " + viper.GetString("bot_token"))
	if botErr != nil {
		panic(fmt.Errorf("Cannot create session: %w", botErr))
	}

	fmt.Println("Preparing file...")

	if err := os.Mkdir(temporaryPath, 0755); os.IsExist(err) {
		os.RemoveAll(temporaryPath)
		os.Mkdir(temporaryPath, 0755)
	}

	runFFmpeg := exec.Command(ffmpegPath,
		"-i", workingFilePath,
		"-codec:", "copy",
		"-start_number", "0",
		"-hls_segment_filename", temporaryPath+"/"+fileNameTemp+"_%03d.ts",
		"-hls_time", viper.GetString("file_max_second"),
		"-hls_list_size", "0",
		"-f", "hls",
		temporaryPath+"/.discord-vugo-main.m3u8")

	if errors.Is(runFFmpeg.Err, exec.ErrDot) {
		runFFmpeg.Err = nil
	}

	if err := runFFmpeg.Run(); err != nil {
		panic(fmt.Errorf("Error during execute ffmpeg command: %w", err))
	}

	tmpFiles, err := os.ReadDir(temporaryPath)
	if err != nil {
		panic(fmt.Errorf("Cannot read temporary directory: %w", err))
	}

	playlistFile, err := os.ReadFile(temporaryPath + "/.discord-vugo-main.m3u8")
	if err != nil {
		panic(fmt.Errorf("Cannot read playlist file in temporary directory: %w", err))
	}

	progressUploading := progressbar.NewOptions(len(tmpFiles)-1,
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("Uploading to Discord..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for _, spilited := range tmpFiles {
		if path.Ext(spilited.Name()) == ".ts" && strings.Contains(spilited.Name(), fileNameTemp) {

			file, err := os.Open(temporaryPath + "/" + spilited.Name())
			if err != nil {
				fmt.Println("Error opening file "+spilited.Name()+":", err)
			}

			res, err := bot.ChannelFileSend(viper.GetString("guild_channel_id"), spilited.Name(), file)
			if err != nil {
				fmt.Println("Error uploading file "+spilited.Name()+":", err)
				return
			}

			fileUrl := res.Attachments[0].URL

			if viper.GetBool("use_proxy") == true {
				fileUrl = viper.GetString("proxy_endpoint") + url.QueryEscape(fileUrl)
			}

			playlistFile = []byte(strings.Replace(string(playlistFile), spilited.Name(), fileUrl, -1))

			progressUploading.Add(1)
			file.Close()
		}
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		panic(fmt.Errorf("Cannot create output file: %w", err))
	}

	if _, err := file.Write(playlistFile); err != nil {
		panic(fmt.Errorf("Error writing to output file: %w", err))
	}

	fmt.Print("\nCleaning up... ")
	os.RemoveAll(temporaryPath)

	fmt.Print("OK!")

	bot.Close()
}
