PROJECT_NAME=github.com/csc13010-student-management
MODULE_NAME=$(name)
MODULE_DIR=internal/$(PROJECT_NAME)/$(MODULE_NAME)
MODULE_INTERFACE=$(shell powershell -Command "'$(name)'.Substring(0,1).ToUpper()+'$(name)'.Substring(1)")

run:
	go run cmd/server/main.go

build:
	go build -o student-management.exe main.go

# test:
# 	go test -v ./tests

d-up:
	docker-compose up

d-up-d:
	docker-compose up -d

d-down:
	docker-compose down

d-down-v:
	docker-compose down -v

# -------------------------------------------------------------------------------------

mockgen-package:
	mockgen -source='internal/$(name)/repository.go' -destination='internal/$(name)/mocks/repository_mock.go' -package=mocks
	mockgen -source='internal/$(name)/usecase.go' -destination='internal/$(name)/mocks/usecase_mock.go' -package=mocks
	
mockgen:
	make mockgen-package name=student
	make mockgen-package name=faculty
	make mockgen-package name=program
	make mockgen-package name=status
	make mockgen-package name=fileprocessor
	make mockgen-package name=notification

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
	make testgen-package name=fileprocessor
	make testgen-package name=notification

# -------------------------------------------------------------------------------------

test:
	go test -cover ./...

test-summary:
	gotestsum --junitfile report.xml ./...

test-lint:
	golangci-lint run ./...

test-ci-check:
	@echo "Running lint..."
	$(MAKE) lint
	@echo "Running tests..."
	$(MAKE) test-summary

NOCACHE ?= 1

TEST_CMD = go test $(if $(filter 1,$(NOCACHE)),-count=1) 

test-unit-package:
	$(TEST_CMD) ./internal/$(name)/repository
	$(TEST_CMD) ./internal/$(name)/usecase
	$(TEST_CMD) ./internal/$(name)/delivery/http

test-unit:
# make test-unit-package name=notification NOCACHE=$(NOCACHE)
	make test-unit-package name=fileprocessor NOCACHE=$(NOCACHE)
	make test-unit-package name=status NOCACHE=$(NOCACHE)
	make test-unit-package name=program NOCACHE=$(NOCACHE)
	make test-unit-package name=faculty NOCACHE=$(NOCACHE)
	make test-unit-package name=student NOCACHE=$(NOCACHE)

COVERAGE_FILE = coverage.out

test-cover-package:
	go test -coverprofile=$(name).cover -covermode=atomic ./internal/$(name)/...
	@if exist $(name).cover ( findstr /V "mode:" $(name).cover >> $(COVERAGE_FILE) & del $(name).cover )

test-cover:
	@if exist $(COVERAGE_FILE) del /F /Q $(COVERAGE_FILE)
	@echo mode: atomic > $(COVERAGE_FILE)
# make test-cover-package name=notification
# make test-cover-package name=fileprocessor
	make test-cover-package name=status
	make test-cover-package name=program
	make test-cover-package name=faculty
	make test-cover-package name=student

test-cover-html:
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html

test-cover-all:
	make test-cover 
	make test-cover-html

# -------------------------------------------------------------------------------------

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

.PHONY: generate test test-summary lint testgen ci-check test-unit test-unit-student test-unit-faculty test-unit-program test-unit-status test-unit-notification mockgen