# grpc-repos

## pb生成
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos\hello.proto
```

## 起動
### server
```
cd cmd/server
go run main.go
```

### client
```
cd cmd/client
go run main.go [name]
```

### request
```
grpcurl -plaintext -d '{"name": "kenta"}' localhost:8080 hello.Greeter/SayHello
```