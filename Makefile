PROJECT  = alfred-time-convert
TESTARGS ?= -v -race -cover

.PHONY: dist
dist: build
	rm -f $(PROJECT).alfredworkflow
	(cd build && zip -r "../$(PROJECT).alfredworkflow" .)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/$(PROJECT) -ldflags="-s -w"
	cp _workflow/* build/

.PHONY: clean
clean:
	go clean
	rm -f $(PROJECT).alfredworkflow
	rm build/*

.PHONY: test
test:
	go test ./... $(TESTARGS)
