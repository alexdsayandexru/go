PHONY: generate-structs
generate-structs:
	mkdir -p pkg/user_v1
	protoc -I api api/user_v1/api.proto --go_out=pkg/user_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative