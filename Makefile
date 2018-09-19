all:
	go build -o bin/freeredirect src/freeredirect.go

install:
	go get github.com/jpsheehan/dotenv
	go get github.com/mattn/go-sqlite3

