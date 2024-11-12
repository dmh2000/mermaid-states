.PHONY: test

all: staticcheck test
	$(MAKE) -C statemachine

test:
	@echo "------test state-gen"
	./test.sh

staticcheck:
	@echo "------staticcheck state-gen"
	staticcheck .
