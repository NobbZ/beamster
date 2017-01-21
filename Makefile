SUBDIRS = $(shell find . -mindepth 2 -name Makefile)

GOBINPATH = $(GOPATH)/bin

GOLEX = $(GOBINPATH)/golex

GIT_HOOKS = $(wildcard .hooks/*)

all: $(SUBDIRS:%=mk_%)

hooks: $(GIT_HOOKS:.hooks/%=.git/hooks/%)

mk_%/Makefile:
	make -C $*

.git/hooks/%: .hooks/%
	cp $< $@