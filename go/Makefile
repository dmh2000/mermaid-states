.PHONY: all header test clean

all: header
	$(MAKE) -s  -C cmd
	$(MAKE) -s  -C pkg
	$(MAKE) -s  -C internal

test: header
	$(MAKE) -s -C cmd test
	$(MAKE) -s -C pkg test
	$(MAKE) -s -C internal test

clean: header
	$(MAKE) -s -C cmd clean
	$(MAKE) -s -C pkg clean
	$(MAKE) -s -C internal clean


header: 
	@echo "================== GO =================="
