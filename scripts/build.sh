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
  OUTPATH="${BINDIR}/${NAME}"
  echo "Building ${NAME}..."
  go build -trimpath -ldflags '-w -s' -o "${OUTPATH}" ./cmd/backup-cli

  if [ "$PACK" = "1" ]; then
    if [ "${GOOS}" = "windows" ]; then
      zip -j "${OUTPATH}.zip" "${OUTPATH}"
    else
      tar -C "${BINDIR}" -zcvf "${OUTPATH}.tar.gz" "${NAME}"
    fi
    rm "${OUTPATH}"
  fi
}

build windows amd64
build linux amd64

# Should work for most rpi we need...
# Armv6 (Raspberry Pi 1/Zero/Zero W/CM)
# Armv7 (Raspberry Pi 2)
# Armv8 (Raspberry Pi 3)
build linux armv7
