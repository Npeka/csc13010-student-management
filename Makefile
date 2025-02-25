run:
	go run cmd/server/main.go

build:
	go build -o student-management.exe main.go

test:
	go test -v ./tests

d-up:
	docker-compose up

d-up-d:
	docker-compose up -d

d-down:
	docker-compose down

d-down-v:
	docker-compose down -v