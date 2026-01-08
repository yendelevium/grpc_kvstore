### grpc_kvstore
A simple in-memory threadsafe kv store implemented via gRPC. Just getting familiar with gRPC

### Setup & Run
Generate proto files
```
$ cd proto
$ protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/main.proto
```

Run server (from root)
```
$ go run server/main.go
```

Run client (from root)
```
$ go run client/main.go
```
