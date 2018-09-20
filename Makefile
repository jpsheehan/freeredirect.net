all:
	go build -o bin/freeredirect ./src/freeredirect.go ./src/database.go ./src/account.go ./src/domain.go

run:
	go run ./src/freeredirect.go ./src/database.go ./src/account.go ./src/domain.go

install:
	go get github.com/jpsheehan/dotenv
	go get github.com/mattn/go-sqlite3

