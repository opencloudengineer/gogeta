#!/bin/bash
echo "Building app for Version : ${VERSION} && Git Commit SHA : ${GITHUB_SHA}"

set -x

export GO111MODULE=on
export CGO_ENABLED=0

go get ./...

go vet ./...

mkdir -p ./bin

LDFLAGS="-s -w -X github.com/opencloudengineer/gogeta/cmd.Version=${VERSION} -X github.com/opencloudengineer/gogeta/cmd.GitCommit=${GITHUB_SHA}"
FLAGS='-a -installsuffix cgo -o'
BUILD='go build -ldflags'

GOOS=windows ${BUILD} "${LDFLAGS}" ${FLAGS} ./bin/gogeta-windows-amd64.exe
GOOS=darwin ${BUILD} "${LDFLAGS}" ${FLAGS} ./bin/gogeta-darwin-amd64
GOOS=linux ${BUILD} "${LDFLAGS}" ${FLAGS} ./bin/gogeta-linux-amd64
GOOS=linux GOARCH=arm64 ${BUILD} "${LDFLAGS}" ${FLAGS} ./bin/gogeta-linux-arm64
