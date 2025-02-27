PROJECT_NAME=github.com/csc13010-student-management
MODULE_NAME=$(name)
MODULE_DIR=internal/$(PROJECT_NAME)/$(MODULE_NAME)
MODULE_INTERFACE=$(shell powershell -Command "'$(name)'.Substring(0,1).ToUpper()+'$(name)'.Substring(1)")

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

# -------------------------------------------------------------------------------------
testgen-package:
	gotests -all -w internal/$(name)/repository/repository.go
	gotests -all -w internal/$(name)/usecase/usecase.go
	gotests -all -w internal/$(name)/delivery/http/handlers.go

testgen:
	make testgen-package name=student
	make testgen-package name=faculty
	make testgen-package name=program
	make testgen-package name=status
	make testgen-package name=fileprocessing
	make testgen-package name=notification
	
# -------------------------------------------------------------------------------------
mockgen-package:
	mockgen -source='internal/$(name)/repository.go' -destination='internal/$(name)/mocks/repository_mock.go' -package=mocks
	mockgen -source='internal/$(name)/usecase.go' -destination='internal/$(name)/mocks/usecase_mock.go' -package=mocks
	
mockgen:
	make mockgen-package name=student
	make mockgen-package name=faculty
	make mockgen-package name=program
	make mockgen-package name=status
	make mockgen-package name=fileprocessing
	make mockgen-package name=notification

test:
	go test -cover ./...

test-summary:
	gotestsum --junitfile report.xml ./...

lint:
	golangci-lint run ./...

ci-check:
	@echo "Running lint..."
	$(MAKE) lint
	@echo "Running tests..."
	$(MAKE) test-summary

unit-test-student:
	go test ./internal/student/repository
	go test ./internal/student/usecase
	go test ./internal/student/delivery/http

unit-test-faculty:
	go test ./internal/faculty/repository
	go test ./internal/faculty/usecase
	go test ./internal/faculty/delivery/http

unit-test-program:
	go test ./internal/program/repository
	go test ./internal/program/usecase
	go test ./internal/program/delivery/http

unit-test-status:
	go test ./internal/status/repository
	go test ./internal/status/usecase
	go test ./internal/status/delivery/http

unit-test-notification:
	go test ./internal/notification/repository
	go test ./internal/notification/usecase
	go test ./internal/notification/delivery/http

unit-test:
	make unit-test-student
	make unit-test-faculty
	make unit-test-program
	make unit-test-status
	make unit-test-notification

generate:
	@if not exist internal\$(name) mkdir internal\$(name)
	@echo package $(name) > internal\$(name)\usecase.go
	@echo. >> internal\$(name)\usecase.go
	@echo type I$(MODULE_INTERFACE)Usecase interface {} >> internal\$(name)\usecase.go
	@echo package $(name) > internal\$(name)\repository.go
	@echo. >> internal\$(name)\repository.go
	@echo type I$(MODULE_INTERFACE)Repository interface {} >> internal\$(name)\repository.go
	@echo package $(name) > internal\$(name)\delivery.go
	@echo. >> internal\$(name)\delivery.go
	@echo type I$(MODULE_INTERFACE)Handlers interface {} >> internal\$(name)\delivery.go

	@if not exist internal\$(name)\delivery\http mkdir internal\$(name)\delivery\http
	@echo package http > internal\$(name)\delivery\http\handlers.go
	@echo. >> internal\$(name)\delivery\http\handlers.go
	@echo import "$(PROJECT_NAME)/internal/$(name)" >> internal\$(name)\delivery\http\handlers.go
	@echo. >> internal\$(name)\delivery\http\handlers.go
	@echo type $(name)Handlers struct {} >> internal\$(name)\delivery\http\handlers.go
	@echo. >> internal\$(name)\delivery\http\handlers.go
	@echo func New$(MODULE_INTERFACE)Handlers() $(name).I$(MODULE_INTERFACE)Handlers { >> internal\$(name)\delivery\http\handlers.go
	@echo 	return ^&$(name)Handlers{} >> internal\$(name)\delivery\http\handlers.go
	@echo } >> internal\$(name)\delivery\http\handlers.go

	@echo package http > internal\$(name)\delivery\http\routes.go
	@echo. >> internal\$(name)\delivery\http\routes.go
	@echo import ( >> internal\$(name)\delivery\http\routes.go
	@echo 	"$(PROJECT_NAME)/internal/$(name)" >> internal\$(name)\delivery\http\routes.go
	@echo 	"github.com/gin-gonic/gin" >> internal\$(name)\delivery\http\routes.go
	@echo ) >> internal\$(name)\delivery\http\routes.go
	@echo. >> internal\$(name)\delivery\http\routes.go
	@echo func Map$(name)Handlers(ftGroup *gin.RouterGroup, h $(name).I$(MODULE_INTERFACE)Handlers) { >> internal\$(name)\delivery\http\routes.go
	@echo } >> internal\$(name)\delivery\http\routes.go

	@if not exist internal\$(name)\repository mkdir internal\$(name)\repository
	@echo package repository > internal\$(name)\repository\repository.go
	@echo. >> internal\$(name)\repository\repository.go
	@echo import "$(PROJECT_NAME)/internal/$(name)" >> internal\$(name)\repository\repository.go
	@echo. >> internal\$(name)\repository\repository.go
	@echo type $(name)Repository struct {} >> internal\$(name)\repository\repository.go
	@echo. >> internal\$(name)\repository\repository.go
	@echo func New$(MODULE_INTERFACE)Repository() $(name).I$(MODULE_INTERFACE)Repository { >> internal\$(name)\repository\repository.go
	@echo 	return ^&$(name)Repository{} >> internal\$(name)\repository\repository.go
	@echo } >> internal\$(name)\repository\repository.go

	@if not exist internal\$(name)\usecase mkdir internal\$(name)\usecase
	@echo package usecase > internal\$(name)\usecase\usecase.go
	@echo. >> internal\$(name)\usecase\usecase.go
	@echo import ( >> internal\$(name)\usecase\usecase.go
	@echo 	"$(PROJECT_NAME)/internal/$(name)" >> internal\$(name)\usecase\usecase.go
	@echo ) >> internal\$(name)\usecase\usecase.go
	@echo. >> internal\$(name)\usecase\usecase.go
	@echo type $(name)Usecase struct { >> internal\$(name)\usecase\usecase.go
	@echo 	fr $(name).I$(MODULE_INTERFACE)Repository >> internal\$(name)\usecase\usecase.go
	@echo } >> internal\$(name)\usecase\usecase.go
	@echo. >> internal\$(name)\usecase\usecase.go
	@echo func New$(MODULE_INTERFACE)Usecase(fr $(name).I$(MODULE_INTERFACE)Repository) $(name).I$(MODULE_INTERFACE)Usecase { >> internal\$(name)\usecase\usecase.go
	@echo 	return ^&$(name)Usecase{fr: fr} >> internal\$(name)\usecase\usecase.go
	@echo } >> internal\$(name)\usecase\usecase.go

	@echo Module $(name) under project $(PROJECT_NAME) has been generated successfully!

.PHONY: generate test test-summary lint testgen ci-check unit-test unit-test-student unit-test-faculty unit-test-program unit-test-status unit-test-notification mockgen