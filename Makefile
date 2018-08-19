APP_NAME := tdur
SRCS := $(shell find . -name \*.go)

.PHONY: build
build: $(SRCS) clean
	go build -o bin/$(APP_NAME) .
	go install

.PHONY: test
test: $(SRCS)
	go test -cover ./...

.PHONY: deploy
deploy: test build
	echo TODO

.PHONY: clean
clean:
	-rm -f bin/*
