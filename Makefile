SHELL := /bin/bash

.PHONY : all
all : clean deps build deploy

.PHONY: deps
deps :
	go get github.com/xxx; go get github.com/yyy

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

