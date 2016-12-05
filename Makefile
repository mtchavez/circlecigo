GOPKGS = \
		golang.org/x/tools/cmd/cover \
		golang.org/x/tools/cmd/godoc \
		golang.org/x/tools/cmd/gorename \
		github.com/golang/lint
default: test

deps:
	@go get -u -v $(GOPKGS)

build: pre_test
	@./script/build

pre_test:
	@./script/pre_test

test: pre_test
	@./script/test

PHONY: deps build pre_test test
