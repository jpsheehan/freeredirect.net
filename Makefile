all:
	go build -o bin/freeredirect ./src/freeredirect.go  ./src/routes.go ./src/server.go ./src/database.go ./src/account.go ./src/domain.go ./src/utils.go

run:
	go run ./src/freeredirect.go ./src/routes.go ./src/server.go ./src/database.go ./src/account.go ./src/domain.go ./src/utils.go

install:
	go get github.com/jpsheehan/dotenv
	go get github.com/mattn/go-sqlite3

