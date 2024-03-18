package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test = struct {
		value string
		want  string
		err   error
	}

	tests := []test{
		{value: "", want: "", err: ErrNoAuthHeaderIncluded},
		{value: "something", want: "", err: ErrNoAuthHeaderIncluded},
		{value: "ApiKey ACTUALY_KEY", err: nil},
	}

	httpHeader := http.Header{}
	for _, tc := range tests {
		httpHeader.Set("Authorization", tc.value)
		got, err := GetAPIKey(httpHeader)
		if err != tc.err {
			t.Fatalf("Expected value: %v with error: %v, got value: %v with error: %v", tc.want, tc.err, got, err)
		}
	}
}
