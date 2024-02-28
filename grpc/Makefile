PHONY: generate-structs
generate-structs:
	mkdir -p pkg/user_v1
	protoc -I api/user_v1 --go_out=plugins=grpc:pkg/user_v1 api/user_v1/api.proto