# Go CLI


## Running

`go build` - compile
`go run main.go` - compile and run without creating a native binary
`go test` - run tests
`go fmt` - format code in the current directory
`go fmt ./...` - format code in the current directory and all subdirectories
`go vet` - static analysis
`go doc` - generate documentation

## Modules

`go mod init <project name>` - init a project

## Dependencies

`go get <package-url>` - download and add to the module a dependency
`go get -u ./...` - update all dependencies to the latest minor or patch
`go get -u=patch ./...` - update all dependencies to the latest patch
`go mod tidy` - clean up dependencies
