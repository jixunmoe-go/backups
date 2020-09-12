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

  NAME="backup-cli-${GOOS}-${GOARCH}"
  OUT_PATH="${BINDIR}/${NAME}"
  echo "Building ${NAME}..."
  go build -trimpath -ldflags '-w -s' -o "${OUT_PATH}${EXT}" ./cmd/backup-cli

  if [ "$PACK" = "1" ]; then
    if [ "${GOOS}" = "windows" ]; then
      zip -j "${OUT_PATH}.zip" "${OUT_PATH}${EXT}"
    else
      tar -C "${BINDIR}" -zcvf "${OUT_PATH}.tar.gz" "${NAME}"
    fi
  fi
}

build windows amd64
build linux amd64

# Should work for most rpi we need...
# arm v6 (Raspberry Pi 1/Zero/Zero W/CM)
# arm v7 (Raspberry Pi 2)
# arm v8 (Raspberry Pi 3)
build linux armv7
