# jwn-it
URL Shortener to learn more about golang!

# Development
## Run
To run locally, run:
```
go run main.go
```
To build locally, run:
```
go build
```

## Run Tests
Run recursively (all packages)
```
go test -v ./...
```

Run test for specific package
```
go test -v ./<package>
```
example:
```
go test -v ./routes
```

## Dependencies
When adding new dependencies/packages, make sure to run:
```
go mod init
go mod tidy
```
This is so GitHub Actions can resolve local dependencies when trying to build