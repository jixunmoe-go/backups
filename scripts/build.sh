#!/usr/bin/env sh

set -e

export CGO_ENABLED=0
if [ -z "${BINDIR}" ]; then
  BINDIR=bin
fi

build() {
  export GOOS="$1"
  export GOARCH="$2"
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

# arm: arm v6 (Raspberry Pi 1/Zero/Zero W/CM)
# other arm arch does not seem to be included by default?
build linux arm
