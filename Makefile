generate-sql-local:
	cd "./migrations" && goose mysql "$(user):$(password)@/sharecycle-local?parseTime=true" create $(name) sql