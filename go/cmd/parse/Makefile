.PHONY: test

all: header staticcheck test

header:
	@echo "------------------ CMD/PARSE ------------------"

test:
	@echo "------test parse"
	@echo " TO TEST, RUN ./test.sh in the terminal"

staticcheck:
	@echo "------staticcheck parse"
	staticcheck .

clean: header
	rm -f *.o
	go clean