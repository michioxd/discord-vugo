# discord-vugo

<ins>**discord**</ins> <ins>**video**</ins> <ins>**uploader**</ins> written in <ins>**g**</ins>olang. A simple way to host your video on Discord via the HLS method.

## Requirements

- A application in the [Discord Developer Portal](https://discord.com/developers/applications)
- [FFmpeg](#install-ffmpeg) installed to your `PATH`

## Usage

- Download binaries from [Release page](https://github.com/michioxd/discord-vugo/releases/latest).
- Open terminal in the folder have `discord-vugo` binary then type `./discord-vugo.exe` (or `chmod +x ./discord-vugo` then `./discord-vugo` in Linux)
- On the first run, it will create the configuration file `discord-vugo-config.yaml`.
- Open the `discord-vugo-config.yaml` file with Notepad (or your favorite IDE like `Visual Studio Code`, `Vim`, `nano`,...).
- You may check out the [Configuration](#configuration) section below.

## Configuration

Use those options in [discord-vugo-config.yaml](discord-vugo-config.yaml.example)

#### `bot_token` [REQUIRED]

**Type**: `string`

**Description:** Bot token, claim it from [Discord Developer Portal](https://discord.com/developers/applications). [How to get bot token](#how-to-get-discord-bot-token)

#### `guild_channel_id` [REQUIRED]

**Type**: `string`

**Description:** Channel ID to upload file into it, right click to a channel you need to get ID, click `Copy Channel ID`. Make sure you add your bot to the server and give your bot permission to send attachments to that channel in your server.

#### `input_file` [REQUIRED]

**Type**: `string`

**Description:** Input file path to upload. Optional if you pass it via `-i` argument.

#### `output_file`

**Type**: `string`

**Default**: `./output.m3u8`

**Description:** Output playlist file path. Can pass it via `-o` argument.

#### `use_proxy`

**Type**: `boolean`

**Default**: `false`

**Description:** Enable proxy URL.

#### `proxy_endpoint`

**Type**: `boolean`

**Description:** Required if `use_proxy` is enabled. Endpoint proxy will write to playlist file. Example with `https://example.com/proxy?url=`, output in the playlist file will be `https://example.com/proxy?url=https%3A%2F%2Fcdn.discordapp.com%2Fattachments%2F1145752139927388313%2F1146790582899978340%2FSuisei_is_Talalala.mp4%3Fex%3D65ce09f5%26is%3D65bb94f5%26hm%3D3ab28624177860467514e24543d7cd34a6a87ed278e06003f679b58052860cde%26`

## Arguments

**Usage:** `discord-vugo <arg> <value>`

**`-i <input file path>`:** Required. Input file path to upload. Optional if you pass it via `input_file` option in configuration file.

**`-o <output file path>`:** Optional. Default `./output.m3u8`. Output playlist file path. Can pass it via `output_file` option in configuration file.

## Tutorial

### How to get Discord Bot Token

- Go to [Discord Developer Portal](https://discord.com/developers/applications)
- Click `New Application`

  ![image](https://github.com/michioxd/discord-vugo/assets/80969068/db5ad43f-f64f-43fc-a06e-a1a4f67a2476)

- Enter a name you want (of course bruh), remember to agree Discord TOS :)

  ![image](https://github.com/michioxd/discord-vugo/assets/80969068/8649070e-61f3-4618-93a6-321f08266442)

- Go to `Bot` at sidebar
- Click `Reset Token`
- Enter 2FA code if you have
- Copy the new token

  ![image](https://github.com/michioxd/discord-vugo/assets/80969068/c972ecb0-3722-48df-ad6f-885a768d111f)

### Install FFmpeg

**Windows**

Via Chocolatey

```batch
choco install ffmpeg
```

**Linux**

```bash
# Ubuntu, Debian,...
sudo apt install ffmpeg

# CentOS
yum install ffmpeg

# Arch Linux
pacman -S ffmpeg
```

## License

[MIT License](LICENSE)

## Credits

[me.](https://github.com/michioxd)
