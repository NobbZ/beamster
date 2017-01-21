SUBDIRS = $(wildcard */.)

GOBINPATH = $(GOPATH)/bin

GOLEX = $(GOBINPATH)/golex

all: scanner.go parser.go

scanner.go: scanner.l
	$(GOLEX) -o $@ $<

parser.go: parser.y
	go tool yacc -v $(@:%.go=%.output) -o $@ $<
