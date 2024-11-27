.PHONY: test

all: staticcheck test

header:
	@echo ------------------ STATEMACHINE ------------------

test:
	@echo "------test statemachine"
	go test -timeout 30s .

staticcheck:
	@echo "------staticcheck statemachine"
	staticcheck .
