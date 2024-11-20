.PHONY: test

all: staticcheck test

test:
	@echo "------test graph"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck graph"
	staticcheck .
