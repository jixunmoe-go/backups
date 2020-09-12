#!/usr/bin/env sh

set -e

export CGO_ENABLED=0
if [ -z "${BINDIR}" ]; then
  BINDIR=bin
fi

build() {
  GOOS="$1"
  GOARCH="$2"
  EXT=""
  if [ "${GOOS}" = "windows" ]; then
    EXT=".exe"
  fi

  NAME="backup-cli-${GOOS}-${GOARCH}${EXT}"
  echo "Building ${NAME}..."
  go build -v -trimpath -ldflags '-w -s' -o "${BINDIR}/${NAME}" ./cmd/backup-cli
}

build windows amd64
build linux amd64

# Should work for most rpi we need...
# Armv6 (Raspberry Pi 1/Zero/Zero W/CM)
# Armv7 (Raspberry Pi 2)
# Armv8 (Raspberry Pi 3)
build linux armv7
