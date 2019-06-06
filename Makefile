.ONESHELL:
CGO_ENABLED := $(or ${CGO_ENABLED},0)                                                                                                                                                                              
GO := go                                                                                                                                                                                                           
GO111MODULE := on

.PHONY: test
test:
	CGO_ENABLED=1 $(GO) test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out

.PHONY: generate-client
generate-client:
	rm -rf api
	mkdir -p api
	GO111MODULE=off swagger generate client -f metal-api.json -t api --skip-validation
