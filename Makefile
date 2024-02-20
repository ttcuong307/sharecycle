generate-sql-local:
	cd "./migrations" && goose mysql "$(user):$(password)@/sharecycle-local?parseTime=true" create $(name) sql

run:
	cd "./cmd/app" && go run main.go

mock-usecase:
	mockgen -source=./internal/usecase/$(usecase).go -destination=./internal/usecase/mock/$(usecase).go -package mock_usecase

mock-repository:
	mockgen -source=./internal/repository/db/$(repo).go -destination=./internal/repository/mock/$(repo).go -package mock_repository

proto-window:
	cd "./internal/pb" && del /f /q *.go
	protoc --proto_path="./internal/proto" --go_out="./internal/pb" --go_opt=paths=source_relative --go-grpc_out="./internal/pb" --go-grpc_opt=paths=source_relative --grpc-gateway_out="./internal/pb" --grpc-gateway_opt=paths=source_relative ./internal/proto/*.proto

proto-mac:
	rm -f ./internal/pb/*.go
	protoc --proto_path="./internal/proto" --go_out="./internal/pb" --go_opt=paths=source_relative \
	--go-grpc_out="./internal/pb" --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out="./internal/pb/rest" --grpc-gateway_opt=paths=source_relative \
	./internal/proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: mock