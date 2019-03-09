# go-post-sample-app

[![CircleCI](https://circleci.com/gh/rema424/go-post-sample-app.svg?style=svg)](https://circleci.com/gh/rema424/go-post-sample-app)

## ローカル開発サーバーでのリモートデバッグ

VS Code 上で Go の拡張機能をインストール。

VS Code 上で [F1] > [Go: Install/Update Tools] > [全項目にチェック] > [OK]

`.vscode/launch.json` に設定を記載。

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "App Engine",
      "type": "go",
      "request": "launch",
      "mode": "remote",
      "remotePath": "${workspaceFolder}",
      "port": 2345,
      "host": "127.0.0.1",
      "program": "${workspaceFolder}",
      "env": {},
      "args": []
    }
  ]
}
```

ターミナルを2つ開いて次のコマンドを実行。

```bash
# ローカルサーバーの起動
goapp serve -debug src/app/main/app.yaml

# デバッガーを起動してサーバーにアタッチ
dlv attach $(ps u | grep _go_ap[p] | awk '{print $2}') --headless --listen=127.0.0.1:2345 --api-version=2
```

VS Code 上でソースコードにブレイクポイントを設置。

VS Code 上でデバッグを開始 (F5) 。

ブラウザでローカルサーバーにアクセス。

ブレイクポイントで止まる。