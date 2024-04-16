go-vendor:
	go mod vendor

go-tidy:
	go mod tidy

tidy: go-vendor go-tidy