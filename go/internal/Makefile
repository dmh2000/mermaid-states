.PHONY: test clean

header: 
	@echo "================== INTERNAL =================="

all: header
	$(MAKE) -s  -C parser
	$(MAKE) -s  -C graph
	$(MAKE) -s  -C build

clean: header
	$(MAKE) -s -C parser clean
	$(MAKE) -s -C graph clean
	$(MAKE) -s -C build clean
