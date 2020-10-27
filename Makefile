.PHONY: gen
gen:
	protoc -I api/ --go_out=plugins=grpc:pkg/ api/api.proto

.PHONY: build
build:
	go build -o ./bin ./cmd/server

.PHONY: run
run:
	./bin/server

.PHONY: test
test:
	go test ./...

# run lintfix
.PHONY: lintfix
lintfix:
    ifeq ($(UNAME), Linux)
		find . \( -path './cmd/*' -or -path './internal/*' -or -path './pkg/*' -or -path './e2e/*' \) \
		-type f -name '*.go' -print0 | \
		xargs -0  sed -i '/import (/,/)/{/^\s*$$/d;}'
    endif
    ifeq ($(UNAME), Darwin)
		find . \( -path './cmd/*' -or -path './internal/*' -or -path './pkg/*' -or -path './e2e/*' \) \
		-type f -name '*.go' -print0 | \
		xargs -0  sed -i '' '/import (/,/)/{/^\s*$$/d;}'
    endif
	goimports -local=gitlab.ozon.ru/bx/marketing-info-api -w ./cmd ./internal