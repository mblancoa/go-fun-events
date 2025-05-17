code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./adapters/*

test:
	go clean -testcache
	go test ./...
