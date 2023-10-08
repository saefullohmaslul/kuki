.PHONY: gen-protoc
gen-protoc:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Generating protobuf files..."

	@protoc -I ./proto --go_out=./ --go-grpc_out=require_unimplemented_servers=false:./ --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./proto/*/*.proto

	@echo "Done"

.PHONY: build
build:
	@go build -o ./dist/kuki ./internal

.PHONY: run
run:
	@./dist/kuki
