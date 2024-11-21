BINARY_NAME=tini-go
GOOS=linux

build:
	GOOS=$(GOOS) go build -o $(BINARY_NAME)


clean:
	rm -f $(BINARY_NAME)


.PHONY: build clean