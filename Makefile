APPS        := server
BLDDIR      ?= bin

default: clean build

build: $(APPS)

$(BLDDIR)/%:
	go build $(LDFLAGS) -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done
