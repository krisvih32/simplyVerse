.PHONY: install

install:
	go install
uninstall:
	go clean -i
ifneq ($(VERBOSE),1)
.SILENT:
endif
