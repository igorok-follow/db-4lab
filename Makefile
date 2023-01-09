GOPATH = $(HOME)/go

GRPC_GATEWAY_INSTALLED_PATH = /Users/$(USER)/go/bin/protoc-gen-grpc-gateway
GRPC_SWAGGER_INSTALLED_PATH = /Users/$(USER)/go/bin/protoc-gen-swagger

.ONESHELL:

.PHONY: gen
gen:
	protoc -I ./api/proto \
				-I /Users/igorok/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--go_out=plugins=grpc:./ \
				--plugin=protoc-gen-grpc-gateway=/Users/igorok/go/bin/protoc-gen-grpc-gateway \
				--plugin=protoc-gen-grpc-gateway=/Users/igorok/go/bin/protoc-gen-grpc-gateway \
				--grpc-gateway_out ./ \
				./api/proto/api/*.proto;