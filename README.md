# portfolio
コード公開前提のポートフォリオ用リポジトリ。

簡易な TODO アプリ用の API を想定する。
なお将来的に機能が複雑化すると仮定して、ドメイン駆動設計の考えを取り入れている。

## Getting Started
### 環境変数を設定
```
cp sample.env .env
```

### Docker 起動
```
docker compose up -d
```

### ローカルサーバー起動
```
go run app/server.go
```