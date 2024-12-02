.PHONY: all headertest clean

all: header
	$(MAKE) -s  -C statemachine

test: header
	$(MAKE) -s -C statemachine test

clean: header
	$(MAKE) -s -C statemachine clean

header: 
	@echo "================== PKG =================="
