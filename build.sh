#!/bin/sh

GOOS=windows GOARCH=amd64 go build -o papercomparepdf-win-amd64
GOOS=darwin GOARCH=amd64 go build -o papercomparepdf-mac-amd64
GOOS=darwin GOARCH=arm64 go build -o papercomparepdf-mac-arm64
GOOS=linux GOARCH=amd64 go build -o papercomparepdf-linux-amd64
GOOS=linux GOARCH=arm64 go build -o papercomparepdf-linux-arm64
