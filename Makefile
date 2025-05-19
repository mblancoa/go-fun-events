code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./adapters/*

test:
	go clean -testcache
	go test ./...

build-api:
	go build -o ./target/userapi ./cmd/userapi.go

build-supply:
	go build -o ./target/supply ./cmd/supply.go

build: build-api build-supply