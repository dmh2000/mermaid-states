

.PHONY: all clean header test


all: header
	$(MAKE) -s  -C go 

clean: header
	$(MAKE) -s -C go clean

test: header
	$(MAKE) -s -C go test

header: 
	@echo "==================" $(shell basename $(PWD)) "=================="
