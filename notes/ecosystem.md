# Go ecosystem

## Modules

A module is a collection of packages.
It is equal to a project in a Scala ecosystem.
A module must have a `go.mod` file, where you specify its code repository, Go version, dependencies etc.

Modules can be of the following types:
* Executable module - has a main package
* Library module - has no main package
* Mixed module - has a main package but also can be published as a library

### Naming

Single word names are prefered, but the `kebab-case` is also common for multipe words names.
Use descriptive but laconic names.
Don't add the `go-` prefix.

### Versioning

Go modules use the server with the `v` prefix.
For instance, `v1.2.3`.

It is not necessery to specify a version of you add as a dependency a module of `v1` version.
However, you must specify a semver prefix for the upcoming major versions, e.g. `v2`.

For instance:
```go
import "github.com/user/myproject"     // v1.x.x
import "github.com/user/myproject/v2"  // v2.x.x  
```

## Libraries

In Go, a library must be published on GitHub or GitLab.
Then the librarary is automatically cached by proxy.golang.org, which is an analogue of Maven in the Go ecosystem.
It happens when somebody runs `go get` for the library.
The proxy works here as a fallback in a case GitHub is down, but also to speed up the library's downloading.

The Go client automatically downloads a library from GitHub or the Proxy,
depending if the library's version is already cached in the Proxy.

## Dependencies

You must avoid direct editing of `go.mod` for adding dependencies.
Instead, you the command `go get ...`.
It will download the dependency and will modify the `go.mod` file.
The direct editing works on a compilation.
But in such the case, you don't validate the dependency, Go doesn't resolve the transitive dependencies,
and an editor can have troubles with recognizing the dependency.

## Compilation

Go requires no runtime separately installed on a machine, comparing to the JVM languages.
A Go application is compiled directly to a native code executed on a CPU.

The resulted binary contains all the components required for a stand-alone running of the application:
* The application's code
* Go standard library
* Third-party libraries
* Go runtime

An important note is that Go has a garbage collector, and it is a part of the runtime which is compiled into the binary.

A Go binary is compiled separately per each OS + architecture combination.
