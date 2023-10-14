ifeq ($(OS), Windows_NT)
	EXT :=.exe
endif

.PHONY: help
help: ## Show help command
	@clear
	@printf "===================================================================\n"
	@printf "\033[1mWelcome to kuki UseCase\033[0m\n";
	@printf "===================================================================\n"
	@grep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: protoc
protoc: ## Generate protobuf files
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@echo "Generating protobuf files..."
	@protoc -I ./proto --go_out=./ --go-grpc_out=require_unimplemented_servers=false:./ --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./proto/todos/*.proto ./proto/google/*.proto
	@echo "Done"

.PHONY: build
build: ## Build server binary
	@go build -o ./dist/kuki${EXT} ./internal

.PHONY: run
run: ## Run server in production mode
	@./dist/kuki

.PHONY: run-dev
run-dev: ## Run server in development mode (will restart if any changes)
	@air -c .air.toml

.PHONY: docker-compose
up: ## Run server in docker-compose
	@docker-compose up -d