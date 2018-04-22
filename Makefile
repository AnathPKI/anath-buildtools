
GOPATH = $(shell pwd)

RUNGO=env GOPATH=$(GOPATH) go

all: tests restart-build trigger-build

trigger-build: src/zhaw.ch/anath/trigger-build/trigger-build.go
	$(RUNGO) build zhaw.ch/anath/trigger-build

restart-build: src/zhaw.ch/anath/restart-build/restart-build.go
	$(RUNGO) build zhaw.ch/anath/restart-build

tests:
	$(RUNGO) test zhaw.ch/anath/travis

clean:
	rm -rf src/golang.org
	rm -rf src/github.com
	rm -rf pkg bin
	rm -f restart-build trigger-build

.PHONY: tests clean