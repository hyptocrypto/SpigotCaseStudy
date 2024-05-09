#!/usr/bin/env bash

rm -rf bin/*

GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64.exe &&
GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 &&
GOOS=darwin GOARCH=arm64 go build -o bin/app-darwin-arm64 &&
GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64

