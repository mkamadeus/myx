build:
	go build -ldflags "-w -s" -o myx cli/main.go
.PHONY: build