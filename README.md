# portfolio
コード公開前提のポートフォリオ用リポジトリ。

簡易な TODO アプリ用の API を想定する。
なお将来的に機能が複雑化すると仮定して、ドメイン駆動設計の考えを取り入れている。

## Getting Started
### 環境変数を設定
```
cp sample.env .env
```

### 必要なパッケージをインストール
```
go install github.com/magefile/mage@v1.15.0
mage install
```

### Docker 起動
```
docker compose up -d
```

### マイグレーション実行（Up, Down）
```
mage migrate:up
mage migrate:down
```

### マイグレーションのスキーマバージョン強制変更
```
mage migrate:force <バージョン番号>
```

### ローカルサーバー起動
```
go run application/server/main.go
```

## その他、参考情報
### ファイル・コードの自動生成
MySQL DB を変更した際に実行する
```
 mage generate:sqlboiler
```
