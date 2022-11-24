
install: ## Install dependent go package
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest
db: ## Create need tables
	@go run main.go init
dep: ## Install requirement package
	@go mod tidy
run: ## Run server
	@go run main.go start