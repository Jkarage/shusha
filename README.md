# vuta

A tool to download files faster, Using concurrency in golang

Downloading using a single thread.

``` bash
make test-single
```

Download a file:
```
go run cmd/vuta/main.go --url https://link.testfile.org/PDF20MB
```

This downloads a 20MB testing file in the root of the project.

IN PROGRESS

- Using multiple threads to download a file
