ROOT=`pwd`
export DB_HOST := 127.0.0.1
export DB_USER := auth
export DB_PASSWORD := P@ssword
export DB_NAME := authdb
export DB_PORT := 8432
export ACCESS_SECRET := MySecret
export SERVER_ADDRESS := 127.0.0.1:8888
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: db
db: ## Build db in docker.
	docker-compose build
	docker-compose up -d

.PHONY: compile
compile: ## Compile the proto file.
	protoc --go_out=pkg/pb --go-grpc_out=pkg/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/messages.proto


.PHONY: server
server: ## Build and run server.
	cd cmd/authservice
	go build -o bin/authserver cmd/authservice/main.go
	bin/authserver


.PHONY: testclient
testclient: ## Build and run client. Run with arguments: make testclient username=user password=P@ssword
	cd cmd/authservice
	pwd
	go build -o bin/client cmd/testclient/main.go
	bin/client -username $(username) -password $(password)

