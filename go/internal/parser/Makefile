

.PHONY: all headertest clean

all: header staticcheck test

test:
	@echo "------test parser"
	go test -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck parser"
	staticcheck .

clean: header
	rm -f *.o
	go clean

header:
	@echo "------------------" $(shell basename $(PWD)) "------------------"
