## go development template

GoとVSCodeで開発を進めるためのテンプレートです。

DBとBackendでGoを立ち上げる想定でDocker Composeを用意しています。
必要に応じてFrontendを追加すると良いかも知れません。

docker compose up で立ち上がります

## 前提条件
以下のインストールが必要です。

* [pre-commit](https://github.com/pre-commit/pre-commit)
* 各種goモジュール
```
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest && \
go install github.com/go-critic/go-critic/cmd/gocritic@latest && \
go install github.com/cosmtrek/air@latest && \
go install github.com/securego/gosec/v2/cmd/gosec@latest && \
go install mvdan.cc/gofumpt@latest && \
go install sourcegraph.com/sqs/goreturns@latest && \
go install mvdan.cc/gofumpt@latest && \
go install github.com/go-delve/delve/cmd/dlv@latest
```

## 変更が必要なところ

1. `backend/go.mod` のmodule名を自分のmodule名に変更
2. トップディレクトリの `.env.example` を `.env` にリネームし、`COMPOSE_PROJECT_NAME` の名前を(1)で設定したmodule名に変更
3. `backend/.golangci.yml` の `module-path` を(1)のmodule名に変更
4. `docker/docker-compose.yml` の `go-develop-template` を(1)のmodule名変更(少し多いので注意)


### VSCodeでデバッグする

`backend/.air.toml`の以下を参照し、適宜対応してください。

```
# デバッグを有効にする場合はこちらのコマンドを使ってください
#full_bin = "APP_ENV=dev APP_USER=air dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./tmp/main"
# デバッグを無効にする場合はこちらのコマンドを使ってください
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
```

### 特徴

Linterのベースに [gofumpt](https://github.com/mvdan/gofumpt) を利用しています。これが起因となりvscodeのsetting.json周りをかなりカスタマイズする必要がありました。
また、Linterにはgolangci-lintを利用しており、（バージョンが古くて使えない・ルールが厳しいなどで）不要そうなものや重複する役割りのモジュールはdisableに追加しています。
