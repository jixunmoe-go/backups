Backups
=======

README: **English** | [简体中文][README_CN]

A set of cli tools to help backup.

Note: Please store private key & run decrypt command from a safe environment
      (e.g. generate & store private key offline).
      [Go _does not_ clear AES key in-memory after initialisation][21865].

The program will try to erase private key from the memory
  in the best effort, but there's no guarantee that it will be. 

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

SSH Shell
 ssh        Enter SSH Shell mode and parse commands from $SSH_ORIGINAL_COMMAND.

SSH Shell
 gen        Generate a private key.
 pubkey     Get public key from a given private key.

Encryption
 encrypt    Encrypt bytes from stdin (with pubkey) and write to stdout.
 decrypt    Decrypt bytes from stdin (with privkey) and write to stdout.

Admin
 clean      Remove old archives.

Backup Management
 save       Save content received from stdin to a specified location.
 load       Load content stored in the backup server.
 list       List backup projects, or versions of a given backup project.
 verify     Verify all or a specific project/version.
```

## Usage

```
$ backup-cli gen | tee priv | backup-cli pubkey > pub
$ cat priv pub
s9OntFH589GQrHrztVDkdM7vwLFk85s1sDP3E7ez9qo=
ZRVTaUl3qLMaBH+KOZWlhsIM52lPNteVx/qe+8pIPFo=
$ echo 'hello!!' | backup-cli encrypt "$(cat pub)" | tee c1.bin | xxd
00000000: 4521 4a58 131e 8f31 4563 3782 dc9f 260b  E!JX...1Ec7...&.
00000010: 58d2 0c14 5798 506e 9986 b754 d853 b43b  X...W.Pn...T.S.;
00000020: f6ab 665a 81ad 5951 ec55 2fbc aca2 723a  ..fZ..YQ.U/...r:
00000030: b03f 516e 98fa d16a 9d1c bf5d 5246 24ca  .?Qn...j...]RF$.
00000040: 4a37 2c9b 3753 00eb 36c0 1bad ae92 3e2a  J7,.7S..6.....>*
00000050: eb3b 3067 ee22 5943 3f32 3835            .;0g."YC?285
$ cat c1.bin | backup-cli decrypt priv | xxd
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

tar zcv /my-backup | ssh backup@example.com -- save backup1-daily
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
tar zcv /my-backup | ssh backup@example.com -- backup-cli save backup1-daily
```

### Add checksum to an existing file

```shell script
sha256sum /path/to/file | head -c 64 | xxd -r -p > /path/to/file.sha256

# To upload with the checksum
cat /path/to/file /path/to/file.sha256 | ssh backup@example.com -- backup-cli save backup1-daily
```

### Duplicate backup

Better to run this inside the backup server. 

```shell script
backup-cli load backup1-daily | backup-cli save backup1-weekly
```

To run it over the network:

```shell script
ssh backup@example.com -- backup-cli load backup1-daily \
  | ssh backup@example.com -- backup-cli save backup1-weekly
```

## TODO

- [ ] Add Unit test & integration test (partial).
- [x] Add command to clean up old backups (e.g. keep last 3)
- [x] Command to load backup `backup-cli load <name> [time="latest"]`
- [x] Command to list backup `backup-cli list [<name>]`
- [x] Some verification mechanism; e.g. generate checksum before the upload.
  - Checksum will be generated on encrypt & verified when decrypting.
  - [x] Add `verify` command (specific project / version or all)
- [x] Pass private key using file instead.

[21865]: https://github.com/golang/go/issues/21865
[README_EN]: https://github.com/jixunmoe-go/backups/blob/master/README.MD
[README_CN]: https://github.com/jixunmoe-go/backups/blob/master/README.zh-CN.MD
