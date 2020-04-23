clean: 
	rm go-sitebuilder
	rm ./docs/chapters/*
.PHONY: clean

build:
	go build
.PHONY: build

