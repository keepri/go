build:
	go build
	
build-run:
	go build && ./rss

run:
	./rss

migrate:
	goose -dir sql/schema sqlite3 ./sql/db/db.sqlite up

migrate-down:
	goose -dir sql/schema sqlite3 ./sql/db/db.sqlite down-to 0

migrate-down-one:
	goose -dir sql/schema sqlite3 ./sql/db/db.sqlite down

start-db:
	turso dev -f ./sql/db/db.sqlite
