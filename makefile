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
	go run cmd/vuta/main.go --url  https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet

help:
	go run cmd/vuta/main.go --help