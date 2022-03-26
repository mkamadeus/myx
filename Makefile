setup:
	@go mod download
.PHONY: setup

build:
	go build -ldflags "-w -s" -o myx main.go
.PHONY: build