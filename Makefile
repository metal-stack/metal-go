METAL_API_VERSION := $(or ${METAL_API_VERSION},$(shell cat VERSION))

release:: generate-client mocks gofmt test;

.PHONY: generate-client
generate-client:
	curl -LO https://raw.githubusercontent.com/metal-stack/metal-api/$(METAL_API_VERSION)/spec/metal-api.json
	$(MAKE) generate-client-local

.PHONY: generate-client-local
generate-client-local:
	yq e -ij ".info.version=\"${METAL_API_VERSION}\"" metal-api.json
	yq e '.info.version' metal-api.json
	rm -rf api
	mkdir -p api
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-v ${PWD}:/work \
		metalstack/builder swagger generate client -f metal-api.json -t api --struct-tags json --struct-tags yaml

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-w /work \
		-v ${PWD}:/work \
		vektra/mockery:v2.26.0 -r --keeptree --inpackage --dir api/client --output test/mocks --all
	go run ./test/client/generate/generate_mock_client.go

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out

.PHONY: golangcicheck
golangcicheck:
	@/bin/bash -c "type -P golangci-lint;" 2>/dev/null || (echo "golangci-lint is required but not available in current PATH. Install: https://github.com/golangci/golangci-lint#install"; exit 1)

.PHONY: lint
lint: golangcicheck
	CGO_ENABLED=1 golangci-lint run
