export GO111MODULE=on

divesites: install-deps
	go build -o divesites ./cmd/divesites/main.go

.PHONY: run
run: divesites
	go run cmd/divesites/main.go

.PHONY: install-deps
install-deps:
	go get

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: test
test: fmt lint
	go test -coverprofile=coverage/coverage.out ./...

