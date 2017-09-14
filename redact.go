package sanitize

import (
	"fmt"
	"net/url"
)

// URI redacts the password in the given uri if a password is present
func URI(uri string) (string, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", fmt.Errorf("failed to parse uri: %s", err)
	}

	// Redact password if present
	if u.User != nil {
		username := u.User.Username()
		_, hasPW := u.User.Password()

		if hasPW {
			u.User = url.UserPassword(username, "REDACTED")
		}
	}

	return u.String(), nil
}

// MustURI is similar to URI but assumes no errors will occur
// If an error occurs it outputs an error string
func MustURI(uri string) string {
	u, err := URI(uri)
	if err != nil {
		return "FAILED_TO_REDACT"
	}

	return u
}
