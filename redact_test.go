package sanitize

import "testing"

func TestURI(t *testing.T) {
	type testCase struct {
		in  string
		out string
	}

	testCases := []testCase{
		// NOOP
		testCase{
			in:  "http://google.com",
			out: "http://google.com",
		},
		// Only user
		testCase{
			in:  "http://user@google.com:1234/path?query=param",
			out: "http://user@google.com:1234/path?query=param",
		},
		// Full case
		testCase{
			in:  "http://user:pass@google.com:1234/path?query=param",
			out: "http://user:REDACTED@google.com:1234/path?query=param",
		},
		// Other scheme
		testCase{
			in:  "postgresql://user:pass@host:5432/db",
			out: "postgresql://user:REDACTED@host:5432/db",
		},
	}

	for _, tc := range testCases {
		u, err := URI(tc.in)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if u != tc.out {
			t.Fatalf("wrong answer, expected %s but got %s", tc.out, u)
		}
	}
}

func TestMustURI(t *testing.T) {
	u := MustURI("")
	if u != "FAILED_TO_REDACT" {
		t.Fatalf("expected FAILED_TO_REDACT msg, got %s", u)
	}
}
