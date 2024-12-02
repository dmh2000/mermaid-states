.PHONY: test clean

all: header staticcheck test

header:
	@echo "------------------ GRAPH ------------------ "

test:
	@echo "------test graph"
	go test -v -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck graph"
	staticcheck .

clean: header
	rm -f *.o
	go clean
