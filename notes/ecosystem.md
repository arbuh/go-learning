# Go ecosystem

## Libraries

In Go, a library must be published on GitHub or GitLab.
Then the librarary is automatically cached by proxy.golang.org, which is an analogue of Maven in the Go ecosystem.
It happens when somebody runs `go get` for the library.
The proxy works here as a fallback in a case GitHub is down, but also to speed up the library's downloading.

The Go client automatically downloads a library from GitHub or the Proxy,
depending if the library's version is alraedy cached in the Proxy.


