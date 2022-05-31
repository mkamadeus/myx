setup:
	@go mod download
.PHONY: setup

build:
	go build -ldflags "-w -s" -o myx main.go
.PHONY: build

examples: clean titanic churn mnist catdog

clean:
	cd examples && ./clean.sh

titanic:
	./myx ./examples/titanic/spec.yaml --output examples/titanic
churn:
	./myx ./examples/churn/spec.yaml --output examples/churn
mnist:
	./myx ./examples/mnist/spec.yaml --output examples/mnist
catdog:
	./myx ./examples/catdog/spec.yaml --output examples/catdog
