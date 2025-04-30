package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "")
	got, err := GetAPIKey(headers)
	want := ErrNoAuthHeaderIncluded
	if err == nil {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyEqual(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey TestKey")
	got, err := GetAPIKey(headers)
	want := "TestKey"
	if err != nil {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyFirstWord(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "WrongFirstWord TestKey")
	got, err := GetAPIKey(headers)
	want := ""
	if err == nil {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyTooShort(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "TestKey")
	got, err := GetAPIKey(headers)
	want := ""
	if err == nil {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
