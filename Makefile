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

mockgen-student:
	mockgen -source='internal/student/repository.go' -destination='internal/student/mocks/repository_mock.go' -package=mocks
	mockgen -source='internal/student/usecase.go' -destination='internal/student/mocks/usecase_mock.go' -package=mocks

test-student:
	go test ./internal/student/repository
	go test ./internal/student/usecase
	go test ./internal/student/delivery/http