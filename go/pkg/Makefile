MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))


.PHONY: all headertest clean

all: header
	$(MAKE) -s  -C statemachine

test: header
	$(MAKE) -s -C statemachine test

clean: header
	$(MAKE) -s -C statemachine clean

header: 
	@echo "==================" $(CURRENT_PATH)
