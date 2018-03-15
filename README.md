Dependency Injector
===================

This is a pretty simple service, just a prototype for managing
dependency injections of http handlers.
Apart http.ResponseWriter and *http.Request params http handler may contain other dependencies.
Currently di supports only interfaces for injection, but under the hood it keep all 
dependencies in a map hence it can be extended easily.

Benchmarks:
```bash
go test -bench=.
goos: darwin
goarch: amd64
pkg: test/di
BenchmarkDI-4           20000000               110 ns/op
PASS
ok      test/di 2.412s
```

Resolving each handler dependency takes about 110 ns depending on number of your dependencies.
