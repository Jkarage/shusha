run:
	go run cmd/vuta/main.go

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

test-single:
	go run cmd/vuta/main.go --url https://link.testfile.org/PDF20MB

help:
	go run cmd/vuta/main.go --help