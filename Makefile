.PHONY: gen-protoc
gen-protoc:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Generating protobuf files..."
	@protoc -I ./internal/apps/grpc --go_out=./internal/apps ./internal/apps/grpc/proto/*.proto
	@protoc -I ./internal/apps/grpc --go-grpc_out=./internal/apps ./internal/apps/grpc/proto/*.proto

	@protoc -I ./ --go_out=./ ./proto/*.proto
	@protoc -I ./ --go-grpc_out=./ ./proto/*.proto

	@echo "Done"

.PHONY: build
build:
	@go build -o ./dist/kuki ./internal/modules

.PHONY: run
run:
	@./dist/kuki