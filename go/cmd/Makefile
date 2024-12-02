MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))
.PHONY: all header test clean


all: header
	$(MAKE) -s  -C parse

test: header
	$(MAKE) -s -C parse test

clean: header
	$(MAKE) -s -C parse clean

header: 
	@echo "------------------" $(CURRENT_PATH)
