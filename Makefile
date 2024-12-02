.PHONY: all clean

header: 
	@echo "================== MERMAID =================="

all: header
	$(MAKE) -s  -C go 

clean: header
	$(MAKE) -s -C go clean