.PHONY: test

all: staticcheck test

test:
	@echo "------test parser"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck parser"
	staticcheck .
