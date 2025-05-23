clean:
	find -type f -name 'mock*.go' -print -delete
	go clean

code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./adapters/*

swagger:
	swag init -dir cmd/userapi/  -parseDependency
mocks:
	mockery

test:
	go clean -testcache
	go test ./...

build-api:
	go build -o ./target/userapi ./cmd/userapi

build-supply:
	go build -o ./target/supply ./cmd/supply

build: build-api build-supply