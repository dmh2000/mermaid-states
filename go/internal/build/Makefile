.PHONY: all header test clean
all: header staticcheck test

test:
	@echo "------test build"
	go test -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck build"

clean: header
	rm -f *.o
	go clean


header:
	@echo "------------------ BUILD ------------------"
