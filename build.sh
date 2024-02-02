#!/bin/bash
rm -rf dist
GOOS=windows GOARCH=amd64 go build -o dist/discord-vugo-win-amd64 -ldflags="-s -w"
GOOS=windows GOARCH=386 go build -o dist/discord-vugo-win-i386 -ldflags="-s -w"
GOOS=linux GOARCH=amd64 go build -o dist/discord-vugo-linux-amd64 -ldflags="-s -w"
GOOS=linux GOARCH=386 go build -o dist/discord-vugo-linux-i386 -ldflags="-s -w"
GOOS=darwin GOARCH=arm64 go build -o dist/discord-vugo-darwin-arm64 -ldflags="-s -w"
