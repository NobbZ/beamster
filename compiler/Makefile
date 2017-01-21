SUBDIRS = $(wildcard */.)

GOBINPATH = $(GOPATH)/bin

GOLEX = $(GOBINPATH)/golex

all: scanner.go

scanner.go: scanner.l
	$(GOLEX) -o $@ $<
