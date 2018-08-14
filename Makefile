APP_NAME := tdur
SRCS := $(shell find . -name \*.go)

.PHONY: build
build: $(SRCS) clean
	go build -o bin/$(APP_NAME) .

.PHONY: test
test: $(SRCS)
	go test

.PHONY: deploy
deploy: test build
	echo TODO

.PHONY: clean
clean:
	-rm -f bin/*
