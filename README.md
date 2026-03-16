  # Login Middleware

Lightweight HTTP middleware for Go that extracts login and registration form data and injects authentication information into the request context.

The middleware parses form values (`email` and `password`) and makes them available to downstream handlers via the request context.

---

# Features

- Simple HTTP middleware for login and registration
- Extracts form credentials automatically
- Stores authentication data in request context
- Compatible with the standard `net/http` package
- Integrates with a custom logger

---

# Installation

Add the package to your project.

```go
go get github.com/eugen252009/pkg_go_login/login
```

---

# Basic Usage

Wrap your HTTP handler with the middleware.

```go
http.Handle("/login", login.Login(myHandler))
http.Handle("/register", login.Register(myHandler))
```

Your handler can then access the authentication information from the request context.

---

# Expected Form Fields

The middleware expects the following form fields in the request:

| Field | Description |
|------|-------------|
| `email` | User email address |
| `password` | User password |

Example request body:

```
email=user@example.com
password=secret
```

---

# Middleware Behavior

1. Parses the incoming request form using `r.ParseForm()`.
2. Extracts `email` and `password` values.
3. Stores them inside an `AuthInfo` structure.
4. Injects the structure into the request context.
5. Passes the request to the next handler.

If form parsing fails, the middleware:

- logs the error
- returns an HTTP error response

---

# Example Handler

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // retrieve auth info from context
    auth := ctx.Value(&login.AuthInfo{}).(*login.AuthInfo)

    fmt.Fprintf(w, "Email: %s", auth.Email)
}
```

---

# Logging

The middleware uses the logger package:

```
github.com/eugen252009/pkg_go_logger/logger
```

Errors during form parsing are logged before returning the HTTP error.

---

# Design Goals

This package focuses on simplicity and composability:

- minimal dependencies
- compatible with standard Go HTTP patterns
- easy integration with existing authentication systems

It is intended as a building block for authentication workflows rather than a complete authentication framework.

