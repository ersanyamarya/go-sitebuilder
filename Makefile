clean: 
	rm *.io
	rm ./docs/chapters/*
.PHONY: clean

build:
	go build
.PHONY: build

