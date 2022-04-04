.PHONY: test
# generate error test code
test:
	@protoc -I . -I ../errors \
		--go_out=paths=source_relative:. \
        --go-xerrors_out=paths=source_relative:. \
		test.proto