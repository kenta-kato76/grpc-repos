# grpc-repos

## pb生成
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/user.proto
```

## 起動
### client
```
go run cmd/client/client.go [name]
```

### Docker起動（ローカル）
```
docker compose up -d
docker ps
```

### request
```
grpcurl -plaintext -d '{"name": "kenta"}' localhost:8080 hello.Greeter/SayHello
grpcurl -plaintext -d '{"name": "kenta", "email":"sample@net.com"}' localhost:8080 user.User/CreateUser
grpcurl -plaintext -d '{"id": "1"}' localhost:8080 user.User/GetUser
```


### Env
```
DB_USER=root
DB_PASSWORD=password
DB_NAME=user_service
DB_HOST=localhost
DB_PORT=3306

```


## 構成
### フォルダの役割と構成の確認
- cmd:
    - cmd/server/main.go: サーバーのエントリーポイント。
    - cmd/client/client.go: クライアントのエントリーポイント。

- controller:
リクエストの受け取りと、ユースケースの実行を担当。
    - server.go: gRPCハンドラーやAPIコントローラーの実装が含まれます。

- domain:
ビジネスロジックの中心。エンティティ、リポジトリインターフェース、ドメインサービスを定義。
    - entity: ドメインオブジェクト（エンティティや値オブジェクト）。
    - repository: リポジトリインターフェース。

- external:
外部サービスや外部APIとの通信を扱います。

- infrastructure:
データベースや外部サービスとの通信、技術的な詳細を管理します。
    - database: データベース接続やリポジトリの実装。
    - external: 外部依存サービスとの連携。

- protos:
Protocol Buffersの定義と生成されたコードが含まれます。

- usecase:
ユースケース層。ドメインオブジェクトを使用してビジネスロジックを実行します。

- util:
ログや共通ライブラリなど、アプリケーション全体で使用されるユーティリティを配置。