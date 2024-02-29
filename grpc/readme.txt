$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH=$PATH:/Users/aasudakov/go/bin

protoc --go_out=pkg/user_v1 --go_opt=paths=source_relative api/user_v1/api.proto