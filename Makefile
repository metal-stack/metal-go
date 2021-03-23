METAL_API_VERSION := $(or ${METAL_API_VERSION},$(shell cat VERSION))

.ONESHELL:
CGO_ENABLED := $(or ${CGO_ENABLED},0)
GO := go
GO111MODULE := on

release:: generate-client gofmt test;

.PHONY: gofmt
gofmt:
	GO111MODULE=off $(GO) fmt ./...

.PHONY: test
test:
	CGO_ENABLED=1 $(GO) test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out

.PHONY: generate-client
generate-client:
	curl -LO https://raw.githubusercontent.com/metal-stack/metal-api/$(METAL_API_VERSION)/spec/metal-api.json
	$(MAKE) generate-client-local

.PHONY: generate-client-local
generate-client-local:
	# writing version into the spec as otherwise it's invalid
	docker run --rm --user $$(id -u):$$(id -g) -it -v $(PWD)/metal-api.json:/metal-api.json mikefarah/yq:4 e '.info.version = "$(METAL_API_VERSION)"' -ji /metal-api.json
	rm -rf api
	mkdir -p api
	GO111MODULE=off docker run -it --user $$(id -u):$$(id -g) --rm -v ${PWD}:/work metalstack/builder swagger generate client -f metal-api.json -t api

.PHONY: golangcicheck
golangcicheck:
	@/bin/bash -c "type -P golangci-lint;" 2>/dev/null || (echo "golangci-lint is required but not available in current PATH. Install: https://github.com/golangci/golangci-lint#install"; exit 1)

.PHONY: lint
lint: golangcicheck
	CGO_ENABLED=1 golangci-lint run
