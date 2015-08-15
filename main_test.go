package main

import (
    "testing"
)

func TestSanitize(t *testing.T) {
    if _, err := sanitize("世界abc"); err == nil {
        t.Errorf("expected %q to not be %q", "error", "nil")
    }
    if str, _ := sanitize("ABC"); str != "abc" {
        t.Errorf("expected %q to be %q", str, "abc")
    }
    if str, _ := sanitize("a bc"); str != "a-bc" {
        t.Errorf("expected %q to be %q", str, "a-bc")
    }
}

func TestRandomTLD(t *testing.T) {
    if tld, _ := randomTLD(); tld != "com" && tld != "org" && tld != "net" {
        t.Errorf("expected %q, %q, %q to include %q", "com", "org", "net", tld)
    }
}
