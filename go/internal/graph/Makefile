.PHONY: test

all: header staticcheck test

header:
	@echo "------------------ GRAPH ------------------ "


test:
	@echo "------test graph"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck graph"
	staticcheck .
