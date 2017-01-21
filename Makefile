SUBDIRS = $(shell find . -mindepth 2 -name Makefile)

GOBINPATH = $(GOPATH)/bin

GOLEX = $(GOBINPATH)/golex

all: $(SUBDIRS:%=mk_%)

mk_%/Makefile:
	make -C $*
