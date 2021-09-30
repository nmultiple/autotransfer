GO = go

.PHONY: all
all:
	$(GO) build .

.PHONY: run
run: all
	./autotransfer

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	$(RM) autotransfer
