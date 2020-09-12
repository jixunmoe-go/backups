Backups
=======

README: **English** | [简体中文][README_CN]

A set of cli tools to help backup.

## Encryption

Standard AES (256-bit) Encryption in OFB mode.

### Encrypted file structure

| name       | bytes | comment                                      |
| :--------- | :---- | :------------------------------------------- |
| header     |  `4`  | `E!JX`, fixed 4 bytes.                       |
| public key |  `32` | A randomly generated `curve25519` public key |
| iv         |  `16` | Randomly generated iv values.                |
| data       |  `?`  | AES Encrypted data until the end.            |
| sha256     |  `64` | The `sha256` hash of all bytes above.        |


## Features

```
$ backup-cli
usage: backup-cli <command> [<args>]
Commands available: 

SSH Execution
 ssh      Use 'backup-cli ssh' as forced command.

Key related
 gen      Generate a private key.
 pubkey   Get public key from a given private key.

Encryption
 encrypt  Encrypt bytes from stdin (with pubkey) and write to stdout.
 decrypt  Decrypt bytes from stdin (with privkey) and write to stdout.

Backup management
 save     Save content received from stdin to a specified location.
 load     Load content stored in the backup server.
 list     List backup projects, or versions of a given backup project.
 verify   Verify all or a specific project/version.
```

## Usage

```
$ backup-cli gen | tee priv | backup-cli pubkey > pub
$ cat priv pub
s9OntFH589GQrHrztVDkdM7vwLFk85s1sDP3E7ez9qo=
ZRVTaUl3qLMaBH+KOZWlhsIM52lPNteVx/qe+8pIPFo=
$ echo 'hello!!' | backup-cli encrypt -pubkey "$(cat pub)" | tee c1.bin | xxd
00000000: 4521 4a58 c69c bf5e 22ae d5a3 f3a1 e1e4  E!JX...^".......
00000010: 6884 9b20 114b 8975 5b59 f822 ffc6 0810  h.. .K.u[Y."....
00000020: 6205 dd0f 7114 0a36 caf4 e9eb a12a 6200  b...q..6.....*b.
00000030: 3cf9 f04f ff33 3fe2 9f46 5ea7 a4d6 1e50  <..O.3?..F^....P
00000040: c631 ea97 1ce3 cb89 97f8 34d1 54eb c9e0  .1........4.T...
00000050: 9f3d 6438 b64d f77a 976e 6f41            .=d8.M.z.noA
$ cat c1.bin | backup-cli decrypt -privkey "$(cat priv)" | xxd
00000000: 6865 6c6c 6f21 210a                      hello!!.
```

### Access over SSH

Note: It is recommended to set `backup-cli` as the shell for backups,
  so the user will not have the actual shell or `sftp` access.

In `~/.ssh/authorized_keys`, one can specify `backup-cli` as a shell
  to access a restricted set of commands:

```text
command="/path/to/backup-cli ssh" ssh-rsa ...
```

Then to upload the archive to a remote site, pipe it via `ssh`:

```shell script
# upload backup to somewhere...

tar zcv /my-backup | ssh backup@example.com -- save -name backup1-daily
```

### Commands allowed when set as SSH Shell

Only limited command will be supported over ssh command
  (parsed via `$SSH_ORIGINAL_COMMAND`), and those are:

* Backup Management
  * `save`
  * `load`
  * `list`
  * `verify`

To access the full command set (not recommend),
  call `backup-cli` directly from a shell over ssh. 

### Invoke from SSH without set as a shell

```shell script
tar zcv /my-backup | ssh backup@example.com -- backup-cli save -name backup1-daily
```

### Add checksum to an existing file

```shell script
sha256sum /path/to/file | head -c 64 | xxd -r -p > /path/to/file.sha256

# To upload with the checksum
cat /path/to/file /path/to/file.sha256 | ssh backup@example.com -- backup-cli save -name backup1-daily
```

## TODO

- [ ] Add Unit test & integration test (partial).
- [ ] Add command to clean up old backups (e.g. keep last 3)
- [x] Command to load backup `backup-cli load -name <name> [-time <time>|"latest"]`
- [x] Command to list backup `backup-cli list [-name <name>]`
- [x] Some verification mechanism; e.g. generate checksum before the upload.
  - Checksum will be generated on encrypt & verified when decrypting.
  - [x] Add `verify` command (specific project / version or all)


[README_EN]: https://github.com/jixunmoe-go/backups/blob/master/README.MD
[README_CN]: https://github.com/jixunmoe-go/backups/blob/master/README.zh-CN.MD
