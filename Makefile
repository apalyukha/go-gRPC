gen:
	protoc --proto_path=proto proto/*.proto --go_out=src/server --go-grpc_out=src/server --experimental_allow_proto3_optional
	protoc --proto_path=proto proto/*.proto --go_out=src/client --go-grpc_out=src/client --experimental_allow_proto3_optional

clean:
	rm -rf src/server/pb/
	rm -rf src/client/pb/

server:
	go run src/server/main.go

client:
	go run src/client/main.go

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

test:
	go test -v ./...
