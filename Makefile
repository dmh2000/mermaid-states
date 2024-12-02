MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))


.PHONY: all clean header test


all: header
	$(MAKE) -s  -C go 

clean: header
	$(MAKE) -s -C go clean

test: header
	$(MAKE) -s -C go test

header: 
	@echo "==================" $(CURRENT_PATH)
