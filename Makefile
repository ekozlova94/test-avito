.PHONY: build
build:
	go build -o ./bin/test-avito.exe cmd/test/main.go

.PHONY: run
run:
	go run cmd/test/main.go

.PHONY: test
test:
	go test ./... -v -count 1