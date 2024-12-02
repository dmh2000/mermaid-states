

.PHONY: test clean

all: header staticcheck test

test:
	@echo "------test graph"
	go test -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck graph"
	staticcheck .

clean: header
	rm -f *.o
	go clean

header:
	@echo "------------------" $(shell basename $(PWD)) "------------------"
