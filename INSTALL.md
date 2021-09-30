# インストール方法

このソフトを動作させるためのサーバの設定、Discord 側の設定、メールサーバの設定の 3 つの設定が必要です。

## Discord

通知したいチャンネル (新たに作成しても OK) の設定を開き、「連携サービス」のタブからウェブフックを追加します。
「ウェブフック URL をコピー」のボタンでコピーできる URL を後で使います。

## メールサーバ

転送したいメールを `Circle` という受信ボックスに振り分けられるように設定します。
参考に今の振り分け条件を書いておくと、メールに「学生団体　各位」という文字列が入っていたら振り分けられるようにしています。
(ここはいろいろ工夫できると思うので頑張ってください (他力本願))。

また、そのメールサーバに SMTP で接続できる方法を調べておきます。
Outlook だと[ここ](https://support.microsoft.com/ja-jp/office/outlook-com-%E3%81%AE-pop-imap-%E3%81%8A%E3%82%88%E3%81%B3-smtp-%E3%81%AE%E8%A8%AD%E5%AE%9A-d088b986-291d-42b8-9564-9c414e2aa040)
に書いてあります。

## サーバ (not Discord)

### 前提条件

Systemd が使える Linux (現時点の構成では WSL で動作させるためには追加の設定が必要です)
である必要があります。また、実用面からサーバは常時起動できる必要があるでしょう。
Cron を使用することも可能ですが、追加の設定が必要です。

### 設定

ソースをコンパイルし、インストールします。

```shellsession
$ go build .
$ sudo install -D autotransfer /opt/autotransfer/autotransfer
```

`etc/autotransfer.service` を編集し、Webhook の URL やメールサーバの設定を書き込みます。
`Environment=` で始まる行が編集すべき行です。
以下がキーごとの説明です。

キー                   | 説明
----------------------|--------------------------------------------------------------------------------------
`MXB_MAIL_SERVER`     | メールサーバのアドレスとポート番号を設定します。(例: `example.com:993`)
`MXB_MAIL_USER`       | メールサーバのログイン名を設定します。
`MXB_MAIL_PASSWORD`   | メールサーバのパスワードを設定します。
`MXB_DISCORD_WEBHOOK` | Discord から貰ったウェブフック URL を入れます。(例: `https://discord.com/api/webhooks/xxxx`)

Systemd の設定ファイルをインストールします。

```shellsession
$ sudo install -m 644 etc/autotransfer.service /usr/local/lib/systemd/system/autotransfer.service
$ sudo install -m 644 etc/autotransfer.timer   /usr/local/lib/systemd/system/autotransfer.timer
```

サービスを有効化します。

```shellsession
$ sudo systemctl enable --now autotransfer.timer
$ sudo systemctl start autotransfer.service
```
