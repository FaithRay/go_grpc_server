1. install protoc-gen-go and protoc-gen-go-grpc

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. adjust path env
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

3. generate protobuf
```
protoc --go_out=. --go-grpc_out=. trade.proto
```