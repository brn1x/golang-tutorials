# Annotations from tutorials


## Comands

Run main module
```
go run . | go run ${MODULE_FILE_PATHNAME}
```

Synchronize module and track new created functions
```
go mod tidy
```

Edit mod file to adapt usability on localhost
```
go mod edit -replace example.com/greetings=../greetings
```

Run tests
```
go test -v
```