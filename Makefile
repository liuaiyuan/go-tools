GOPATH:=$(shell go env GOPATH)
.PHONY: gen
gen:
	go generate ./...

.PHONY: test
test:
	go test ./test/...