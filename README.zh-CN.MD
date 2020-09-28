Backups
=======

README: [English][README_EN] | **简体中文**

一套用于辅助备份的命令行工具。

注意：请在安全的环境储存私钥及进行解密过程。
     （如在不同网络的机器上生成并储存私钥）
     [Go _并不会_ 在初始化后清理密钥][21865]。

该程序会尽量清理内存里的私钥内容，但不能提供任何保证。

## 加密

标准 AES (256-bit) 加密，OFB 模式。

### 加密文件结构

| name       | bytes | comment                                      |
| :--------- | :---- | :------------------------------------------- |
| header     |  `4`  | `E!JX`, 固定的 4 字节内容。                     |
| public key |  `32` | 一个随机生成的 `curve25519` 公钥。               |
| iv         |  `16` | 随机生成的 IV （初始向量）值。                    |
| data       |  `?`  | AES 加密后的数据，直到结尾。                      |
| sha256     |  `64` | 从第一个字节直到该域前的所有字节内容的 `sha256` 哈希。|

## 特性

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

## 使用示例

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

### 通过 SSH 访问

注意：推荐将 `backup-cli` 设定为连接的 shell；这样用户就不能获得
  任意代码执行权限以及 `sftp` 权限。

在 `~/.ssh/authorized_keys`, 管理员可以将 `backup-cli` 设定为强制的 shell:

```text
command="/path/to/backup-cli ssh" ssh-rsa ...
```

要上传备份包到远端服务器，通过管道经由 `ssh` 上传： 

```shell script
# 上传备份到某处…

tar zcv /my-backup | ssh backup@example.com -- save backup1-daily
```

### 作为 shell 时可用的指令

当作为 shell 模式运行，只有一部分的指令被允许使用
  （通过环境变量 `$SSH_ORIGINAL_COMMAND` 解析）：

* 备份管理
  * `save` - 储存
  * `load` - 读取
  * `list` - 列出
  * `verify` - 验证备份包签名

如果需要访问完整的指令（不推荐），直接在不限制的 `ssh` shell
  环境执行 `backup-cli` 即可。

### 从 SSH Shell 调用

```shell script
tar zcv /my-backup | ssh backup@example.com -- backup-cli save backup1-daily
```

### 手动计算签名

```shell script
sha256sum /path/to/file | head -c 64 | xxd -r -p > /path/to/file.sha256

# 上传签名后的文件
cat /path/to/file /path/to/file.sha256 | ssh backup@example.com -- backup-cli save backup1-daily
```

## TODO

参见英文说明。


[21865]: https://github.com/golang/go/issues/21865
[README_EN]: https://github.com/jixunmoe-go/backups/blob/master/README.MD
[README_CN]: https://github.com/jixunmoe-go/backups/blob/master/README.zh-CN.MD
