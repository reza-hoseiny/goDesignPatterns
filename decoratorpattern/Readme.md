Decorator pattern is effective to dynamically enhance (decorating) a function by creating a (or several) subclass(es).
By subclassing, we can extend the functionality in a flexible way at runtime.
we can use decorator pattern for writing middleware in Go
Wrappers have the following signature:
func(http.Handler) http.Handler
This approach is used to
Logging and tracing
Validating the request; such as checking authentication credentials
Writing common response headers

A wrapper is an object that can be linked with some target object.
The wrapper contains the same set of methods as the target and delegates to it all requests it receives.
However, the wrapper may alter the result by doing something either before or after it passes the request to the target.
