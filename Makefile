GO = go

.PHONY: all
all:
	$(GO) build .

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	$(RM) autotransfer
