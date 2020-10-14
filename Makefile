.PHONY: gen
gen:
	protoc -I api/ --go_out=plugins=grpc:pkg/ api/api.proto

.PHONY: build
build:
	go build -o ./bin ./cmd/server

.PHONY: run
run:
	./bin/server