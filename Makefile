code-generation:
	find -type f -name '*_impl.go' -print -delete
	go generate ./adapters/*
