.PHONY: test clean

all: header staticcheck test

header:
	@echo "------------------ PARSER ------------------"
	       

test:
	@echo "------test parser"
	go test -v -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck parser"
	staticcheck .

clean: header
	rm -f *.o
	go clean
