.PHONY: test

all: staticcheck test

test:
	@echo "------test statemachine"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck statemachine"
	staticcheck .
