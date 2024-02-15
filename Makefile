generate-sql-local:
	cd "./migrations" && goose mysql "$(user):$(password)@/sharecycle-local?parseTime=true" create $(name) sql

run:
	cd "./cmd/app" && go run main.go