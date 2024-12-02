MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_PATH := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))


.PHONY: test clean

all: header staticcheck test

test:
	@echo "------test statemachine"
	go test -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck statemachine"
	staticcheck .

clean: header
	rm -f *.o
	go clean

header:
	@echo ------------------ $(CURRENT_PATH)
