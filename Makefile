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

mockgen:
	mockgen -source='internal/repository/student_repository.go' -destination='internal/repository/mock/student_repository_mock.go' -package=mock

test-student:
	go test ./internal/student/repository
	go test ./internal/student/usecase
	go test ./internal/student/delivery/http