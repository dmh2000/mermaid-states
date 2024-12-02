MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))


.PHONY: all header test clean
all: header staticcheck test

test:
	@echo "------test build"
	go test -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck build"

clean: header
	rm -f *.o
	go clean


header:
	@echo "------------------ $(CURRENT_PATH)"
