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
 gen      Generate a private key.
 pubkey   Get public key from a given private key.
 encrypt  Encrypt bytes from stdin (with pubkey) and write to stdout.
 decrypt  Decrypt bytes from stdin (with privkey) and write to stdout.
 save     Save content received from stdin to a specified location.
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
$ # now upload to somewhere...
$ tar zcv /my-backup | ssh backup@example.com -- 'backup-cli save -name b1/daily'
```

## TODO

- [ ] Add Unit test & integration test.
- [ ] Add command to clean up old backups (e.g. keep last 3)
- [ ] Command to load backup `backup-cli load -name <name> [-time <time>|"latest"]`
- [ ] Command to list backup `backup-cli list [-name <name>]`
- [ ] Some verification mechanism; e.g. generate checksum before the upload.
