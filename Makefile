.DEFAULT_GOAL := run 

.PHONY: fmt vet run build
fmt:
	go fmt ./...

vet: fmt 
	go vet ./...

run: vet
	go run ./sindhu.go

build: vet
	go build -o bin/sindhu