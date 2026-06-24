PHONY:
GO111MODULE=on

default:

run:
	go run ./cmd/book-club

fmt:
	@gofmt -s -w $$(go list -f "{{.Dir}}" ./...)

gci:
	@gci write -s standard -s default -s "prefix(github.com/Michela-DC/book-club/)" -s blank -s dot ./cmd ./internal

lint-all:
	@golangci-lint run --timeout 2m0s ./...

install-tools:
	@echo Installing tools
	@go install github.com/daixiang0/gci@latest
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	@echo Installation completed