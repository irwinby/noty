.DEFAULT_GOAL := local-run

BINARY_NAME := noty

.PHONY: lint
lint:
	golangci-lint run -v --allow-parallel-runners

.PHONY: clean
clean:
	go clean

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -covermode=atomic -v -race ./...

.PHONY: build
build:
	go build -o ${BINARY_NAME} main.go

.PHONY: build-docker-image
build-docker-image:
	docker build --platform=linux/amd64 -t $(BINARY_NAME):latest .

.PHONY: local-run
local-run: build-docker-image
	docker-compose up --force-recreate --remove-orphans

.PHONY: swagger
swagger:
	swag init --output api/

.PHONY: mock
mock:
	mockery --all --with-expecter --case=underscore --outpkg=mock --output=./internal/mock
