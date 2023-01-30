# Annotations from tutorials


## Comands

Create module/package
```
go mod init example.com/greetings | TEAM_DOMAIN/PROJECT_NAME
```

Run main module
```
go run . | go run ${MAIN_MODULE_FILE_PATHNAME}
```

Synchronize module and track new created functions
```
go mod tidy
```

Install module depencies that are in the package
```
go get .
```

Edit mod file to adapt usability on localhost
```
go mod edit -replace example.com/greetings=../greetings
```

Run tests
```
go test -v
```
#
######Build and install binaries

Generate binary
```
go build
```

Get install directory
```
go list -f `{{.Target}}`
```

Update **PATH** variable
```
export PATH=$PATH:/BIN_DIRECTORY_FROM_LAST_COMMAND_RESULT
```

And finally, to expose the application as a cli command
```
go install
```