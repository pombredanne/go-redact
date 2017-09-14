go-redact
---

Redact strings containing passwords.

### Usage

_URI_

```go
// With error handling
u, err := redact.URI("http://user:pass@github.com")
if err != nil {
  log.Fatalf("failed to redact: %s", err) // Will happen if input is not a valid URI
}

fmt.Println(u)
// Output:
// http://user:REDACTED@github.com
```

```go
// Without error handling
u := redact.MustURI("adkljsaldjksla")

fmt.Println(u)
// Output:
// FAILED_TO_REDACT
```
