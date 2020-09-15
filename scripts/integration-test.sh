#!/usr/bin/env sh

set -e

# Create a native build
go build ./cmd/backup-cli

rm -rf tmp
mkdir -p tmp

export BACKUP_BASE="$PWD/tmp/backups"

PATH="$PATH:$PWD"

echo Generate keys...
backup-cli gen | tee tmp/priv | backup-cli pubkey > tmp/pub
cat tmp/priv tmp/pub

dd if=/dev/urandom bs=4096 count=$((1024 * 4)) | tee tmp/8M \
  | backup-cli encrypt "$(cat tmp/pub)" \
  | backup-cli save test1

backup-cli verify test1
backup-cli load test1 latest | backup-cli decrypt tmp/priv > tmp/decrypted

EXPECTED="$(< tmp/8M sha256sum)"
ACTUAL="$(< tmp/decrypted sha256sum)"

# Clean up
rm -rf tmp backup-cli

if [ "$EXPECTED" = "$ACTUAL" ]; then
  echo "TEST OK!"
else
  echo "TEST FAIL - failed to decrypt original content."
  sha256sum tmp/decrypted tmp/8M
  exit 1
fi
