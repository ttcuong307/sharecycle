generate-sql-local:
	cd "./migrations" && goose mysql "$(user):$(password)@/sharecycle-local?parseTime=true" create $(name) sql

run:
	cd "./cmd/app" && go run main.go

mock-usecase:
	mockgen -source=./internal/usecase/$(usecase).go -destination=./internal/usecase/mock/$(usecase).go -package mock_usecase
.PHONY: mock

mock-repository:
	mockgen -source=./internal/repository/db/$(repo).go -destination=./internal/repository/mock/$(repo).go -package mock_repository
.PHONY: mock