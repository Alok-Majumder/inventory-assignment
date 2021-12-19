GIT_REPO := $(shell git config --get remote.origin.url)
REPO_NAME := $(shell basename ${GIT_REPO} .git)
GOLANGCI_LINT_VERSION := v1.35.2
GOLANGCI_LINT := bin/golangci-lint_$(GOLANGCI_LINT_VERSION)/golangci-lint
GOLINTCI_URL := https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh


## build: Build the application artifacts
build:
	@echo " Creating Binary in ./bin/"${REPO_NAME}
	@go build -o ./bin/${REPO_NAME}

${GOLANGCI_LINT}:
	@mkdir -p $(dir ${GOLANGCI_LINT})
	@curl -sfL ${GOLINTCI_URL} | sh -s -- -b ./$(patsubst %/,%,$(dir ${GOLANGCI_LINT})) ${GOLANGCI_LINT_VERSION}

## lint: Lint the source code
lint: ${GOLANGCI_LINT}
	@echo " Linting code"
	@$(GOLANGCI_LINT) run

## lint-info: Returns information about the current linter being used
lint-info:
	@echo ${GOLANGCI_LINT}

## test: Run Go tests
test:
	@echo " Running tests"
	@go test ./...

## run: Run the application
run: build
	@echo " Running binary"
	@./bin/${REPO_NAME}
