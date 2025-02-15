run:
	go run main.go

build:
	go build -o student-management.exe main.go

test:
	go test -v ./tests