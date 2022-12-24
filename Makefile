NAME=github.com/gurbaj5124871/url-shortener
VERSION=0.0.1
BUF_VERSION:=v1.9.0
export HOST_IP := $(shell ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" | grep -v 127.0.0.1 | awk '{ print $$2 }' | cut -f2 -d: | head -n1)
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

generate: generate/proto
generate/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

lint:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) lint
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) breaking --against 'https://github.com/gurbaj5124871/url-shortener.git#branch=master'

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: build_container
build_container:
	@docker compose build $(call args)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

.PHONY: run_container
run_container:
	@docker compose up $(call args)

.PHONY: restart_container
restart_container:
	@docker compose restart $(call args)

.PHONY: stop_container
stop_container:
	@docker compose stop $(call args)

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -v ./tests/*

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
