.PHONY: test clean

all: header staticcheck test

header:
	@echo ------------------ PKG/STATEMACHINE ------------------

test:
	@echo "------test statemachine"
	go test -v -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck statemachine"
	staticcheck .

clean: header
	rm -f *.o
	go clean
