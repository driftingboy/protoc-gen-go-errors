.PHONY: init
# init dependcy
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

.PHONY: build
# build protoc-gen-go-xerrors
build:
	go build -o protoc-gen-go-xerrors *.go && mv protoc-gen-go-xerrors ${GOPATH}/bin/

.PHONY: generate
# generate error test code
generate:
	@cd ./gerr && protoc -I . \
		--go_out=paths=source_relative:. \
		errors.proto

.PHONY: test
# generate error test code
test:
	@cd ./test && protoc -I . -I ../gerr \
		--go_out=paths=source_relative:. \
        --go-xerrors_out=paths=source_relative:. \
		test.proto && \
	go test ./...