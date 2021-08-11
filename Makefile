# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.


GOBIN = ./build/bin
GO ?= latest
GORUN = go run
GOBUILD = go build

all:
	cd ./src&&$(GOBUILD) main.go&&./main
