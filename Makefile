include .env
export

.PHONY: first
first: ## Build db in docker.
	docker-compose build
	docker-compose up -d

.PHONY: compile
compile: ## Compile the proto file.
	export PATH=$(PATH):$(GOPATH)/bin
	protoc --go_out=pkg/pb --go-grpc_out=pkg/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/messages.proto

.PHONY: server
server: ## Build and run server.
	cd cmd/authservice
	go build -o bin/authserver cmd/authservice/main.go
	bin/authserver


.PHONY: testclient
testclient:  ## Build and run client. Run with arguments: make testclient username=user password=P@ssword
	cd cmd/authservice
	pwd
	go build -o bin/client cmd/testclient/main.go
	bin/client -username $(username) -password $(password)

.PHONY: test  ## Run tests
test:
	go test ./...

.PHONY: m_create  ## Create migration with name
m_create:
	migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: m_up  ## Apply migrations
m_up:
	migrate -source file://db/migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/authdb?sslmode=disable up
