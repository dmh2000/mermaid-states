.PHONY: test

all: header staticcheck test

header:
	@echo "------------------ PARSER ------------------"
	       


test:
	@echo "------test parser"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck parser"
	staticcheck .
