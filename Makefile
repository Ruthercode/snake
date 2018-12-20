commit := $(shell git rev-parse --short HEAD)

all: fmt build

fmt:
	go fmt ./...

build:
	echo $(commit)
	go build ./...

run:
	./edu

clear:
	rm ./edu