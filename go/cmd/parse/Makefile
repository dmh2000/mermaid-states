MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))
# MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
# CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))
.PHONY: all header staticcheck test

all: header staticcheck test

test: header
	@echo "------test parse"
	# @echo " to view test results, run ./test.sh. in the terminal"
	./test.sh > /dev/null 2>&1
	echo "all tests passed"

staticcheck:
	@echo "------staticcheck parse"
	staticcheck .

clean: header
	rm -f *.o
	go clean


header:
	@echo "------------------" $(shell basename $(PWD)) "------------------"


