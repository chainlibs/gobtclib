SHELL := /bin/bash

.PHONY : all
all : clean deps deploy

.PHONY: deps
deps :
	go get -u go.uber.org/zap;go get -u github.com/gobasis/log;

.PHONY: build
build :
	go build -o bbb main.go

.PHONY: clean
clean :
	rm -rf /tmp/mktarget && rm -rf ./target

.PHONY: deploy
deploy : clean
	mkdir ./target

.PHONY: tar
tar :
	cd ./target; tar -zvcf bbb.tar.gz *

