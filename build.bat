@echo off
title discord-vugo builder

rmdir /s /q dist

REM build for windows/amd64
echo windows/amd64
set GOOS=windows
set GOARCH=amd64
go build -o dist/discord-vugo-win-amd64.exe  -ldflags="-s -w"

REM build for windows/i386
echo windows/i386
set GOOS=windows
set GOARCH=386
go build -o dist/discord-vugo-win-i386.exe  -ldflags="-s -w"

REM build for linux/amd64
echo linux/amd64
set GOOS=linux
set GOARCH=amd64
go build -o dist/discord-vugo-linux-amd64  -ldflags="-s -w"

REM build for linux/i386
echo linux/i386
set GOOS=linux
set GOARCH=386
go build -o dist/discord-vugo-linux-i386  -ldflags="-s -w"

REM build for darwin/arm64
echo darwin/arm64
set GOOS=darwin
set GOARCH=arm64
go build -o dist/discord-vugo-darwin-arm64  -ldflags="-s -w"
