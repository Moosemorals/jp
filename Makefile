
EXECUTABLE := jp 

BUILDDIR := ./target

all: test build run

test:
	go test -v

build:
	rm -rf $(BUILDDIR)
	mkdir -p $(BUILDDIR) 
	go build -o $(BUILDDIR)/$(EXECUTABLE)

run: build
	$(BUILDDIR)/$(EXECUTABLE)

clean:
	go clean
	rm -rf $(BUILDDIR) 

