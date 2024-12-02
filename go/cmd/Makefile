.PHONY: all header test clean


all: header
	$(MAKE) -s  -C parse

test: header
	$(MAKE) -s -C parse test

clean: header
	$(MAKE) -s -C parse clean

header: 
	@echo "------------------" $(CURRENT_PATH) "------------------"
