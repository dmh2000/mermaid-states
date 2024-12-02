all: header staticcheck test

header:
	@echo "------------------ BUILD ------------------"

test:
	@echo "------test build"
	go test -v -count=1 -timeout 30s .

staticcheck:
	@echo "------staticcheck build"

clean: header
	rm -f *.o
	go clean