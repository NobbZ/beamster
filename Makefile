SUBDIRS = $(shell find . -mindepth 2 -name Makefile)

GOBINPATH = $(GOPATH)/bin

GOLEX = $(GOBINPATH)/golex
GOMETALINTER = $(GOBINPATH)/gometalinter

GIT_HOOKS = $(wildcard .hooks/*)

all: $(SUBDIRS:%=mk_%) format compile test bench

compile:
	go build ./...
.PHONY: compile

format:
	go fmt ./...

test:
	go test ./...

bench:
	go test ./... -bench .

lint:
	$(GOMETALINTER) --enable-all -e "compiler/(scanner|parser).go" ./... --deadline 30s

hooks: $(GIT_HOOKS:.hooks/%=.git/hooks/%)

mk_%/Makefile:
	make -C $*

.git/hooks/%: .hooks/%
	cp $< $@
	chmod +x $@

golex:
	go get -u github.com/cznic/golex

gometalinter:
	go get -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install
