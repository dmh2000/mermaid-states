.PHONY: test clean

header: 
	@echo "================== GO =================="

all: header
	$(MAKE) -s  -C cmd
	$(MAKE) -s  -C pkg
	$(MAKE) -s  -C internal

clean: header
	$(MAKE) -s -C cmd clean
	$(MAKE) -s -C pkg clean
	$(MAKE) -s -C internal clean
