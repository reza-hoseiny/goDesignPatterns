// we can use decorator pattern for writing middleware in Go
// Wrappers have the following signature:
//  func(http.Handler) http.Handler
// This approach is used to
// Logging and tracing
// Validating the request; such as checking authentication credentials
// Writing common response headers