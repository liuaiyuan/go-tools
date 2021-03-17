GOPATH:=$(shell go env GOPATH)
.PHONY: gen
gen:
	go generate ./...

