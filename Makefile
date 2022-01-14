lint:
	golangci-lint run --build-tags test ./...

# https://github.com/golangci/golangci-lint
install-golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

# https://github.com/psampaz/go-mod-outdated
outdated-list:
	go list -u -m -json all | go-mod-outdated -update -direct
