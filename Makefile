clean:
	find -type f -name 'mock*.go' -print -delete
	go clean

code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./adapters/*

swagger:
	swag init -g cmd/userapi/main.go  -parseDependency
mocks:
	mockery

test:
	go clean -testcache
	go test ./...

build-api:
	rm -f ./bin/userapi
	go build -o ./bin/userapi ./cmd/userapi

build-supply:
	rm -f ./bin/supply
	go build -o ./bin/supply ./cmd/supply

build: build-api build-supply

.PHONY: deploy
deploy:
	docker-compose -f ./deploy/docker-compose.yml up -d

stop:
	docker-compose -f ./deploy/docker-compose.yml down
