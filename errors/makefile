.PHONY: generate
# generate error test code
generate:
	@protoc -I . \
		--go_out=paths=source_relative:. \
		errors.proto