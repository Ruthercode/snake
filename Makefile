commit := $(shell git rev-parse --short HEAD)

all: fmt build

fmt:
	go fmt main.go

build:
	echo $(commit)
	go build main.go

run:
	./main

clear:
	rm ./main