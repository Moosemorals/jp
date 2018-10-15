
EXECUTABLE := jp 

BUILDDIR := ./target

all: test build

test:
	go test -v

build:
	rm -rf $(BUILDDIR)
	mkdir -p $(BUILDDIR) 
	go build -o $(BUILDDIR)/$(EXECUTABLE)

install: test build
	go install 

clean:
	go clean
	rm -rf $(BUILDDIR) 

