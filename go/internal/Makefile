MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))


.PHONY: all header test clean


all: header
	$(MAKE) -s  -C parser
	$(MAKE) -s  -C graph
	$(MAKE) -s  -C build

test: header
	$(MAKE) -s -C parser test
	$(MAKE) -s -C graph test
	$(MAKE) -s -C build test

clean: header
	$(MAKE) -s -C parser clean
	$(MAKE) -s -C graph clean
	$(MAKE) -s -C build clean

header: 
			@echo "==================" $(CURRENT_PATH)
