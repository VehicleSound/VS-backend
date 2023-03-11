.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: build
build:
	go build -v -o "./bin/service" ./cmd


.DEFAULT_GOAL := build