ROOT := $(CURDIR)
PACKAGES := `go list ./... | grep -v /node_modules/`

default: ci

deps:
	@echo "[Deps] getting dependencies"
	GO111MODULE=on go get -u golang.org/x/lint/golint
	GO111MODULE=on go get -u -v ./...

lint:
	@echo "[Lint] running golint"
	@go fmt ${PACKAGES}
	@golint -set_exit_status ${PACKAGES} || exit 1

lint:
	@echo "[Lint] running golint"
	@go fmt ${PACKAGES}
	@golint -set_exit_status ${PACKAGES} || exit 1

vet:
	@echo "[Vet] running go vet"
	go vet ${PACKAGES} || exit 1

ci: deps vet lint test

build: pre_test
	@./script/build

pre_test:
	@./script/pre_test

test: pre_test
	@./script/test

PHONY: ci deps build pre_test test
