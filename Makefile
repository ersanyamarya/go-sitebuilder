all: clean build run pre bundle

clean: 
	rm -f go-sitebuilder
	rm -f ./docs/chapters/*
	rm -fr ./docs/assets
	rm -fr ./docs/*.js
.PHONY: clean

build:
	go build
.PHONY: build

pre:
	cp -rf ./src/assets ./docs
.PHONY: pre 

run: go-sitebuilder
	./go-sitebuilder

runserver:
	npm run start:dev -prefix ./bundler

bundle:
	npm run build -prefix ./bundler
.PHONY: run