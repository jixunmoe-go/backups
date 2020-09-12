Backups
=======

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
00000000: 4521 4a58 4490 5019 33a5 e421 3067 6f3e  E!JXD.P.3..!0go>
00000010: 5487 29b6 2b20 fe71 9403 6a7f 0865 4e57  T.).+ .q..j..eNW
00000020: 9f88 183a af84 97ca ee68 25c1 a111 5899  ...:.....h%...X.
00000030: 464e 6c89 b42a 65e0 4760 8624            FNl..*e.G`.$
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
