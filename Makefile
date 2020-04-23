all: clean build run

clean: 
	rm -f go-sitebuilder
	rm -f ./docs/chapters/*
.PHONY: clean

build:
	go build
.PHONY: build

run: go-sitebuilder
	./go-sitebuilder
.PHONY: run