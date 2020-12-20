#!/bin/bash
set -xeuo pipefail

echo "Building on platform with OS: $(go env GOOS) & Arch: $(go env GOARCH)"
echo "Building app for Version : ${VERSION:=$(git describe --tags)} && Git Commit SHA : ${GITHUB_SHA:=$(git rev-parse HEAD)}"

export GO111MODULE=on
export CGO_ENABLED=0

go get ./...
go vet ./...

mkdir -p ./bin

LDFLAGS="-s -w -X 'github.com/opencloudengineer/gogeta/cmd.Version=${VERSION}' -X 'github.com/opencloudengineer/gogeta/cmd.GitCommit=${GITHUB_SHA}'"

function buildFor() {
  GOOS="${1}" GOARCH="${2}" go build -ldflags="${LDFLAGS}" -a -installsuffix cgo -o "./bin/gogeta-${1}-${2}"
}

buildFor windows amd64
buildFor darwin amd64
buildFor linux amd64
buildFor linux arm64

# add .exe extension to windows binary
mv ./bin/gogeta-windows-amd64{,.exe}
