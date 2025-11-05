package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		inputHeader http.Header
		wantString  string
		wantError   error
	}{
		{
			name: "Correct",
			inputHeader: http.Header{
				"Authorization": {"ApiKey Atoken"},
			},
			wantString: "Atoken",
			wantError:  nil,
		},
		{
			name: "Empty Auth header",
			inputHeader: http.Header{
				"Authorization": {},
			},
			wantString: "",
			wantError:  ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed header",
			inputHeader: http.Header{
				"Authorization": {"Malformation is I"},
			},
			wantString: "",
			wantError:  errors.New("malformed authorization header"),
		},
		{
			name: "Many splits",
			inputHeader: http.Header{
				"Authorization": {"ApiKey To many splits"},
			},
			wantString: "To",
			wantError:  nil,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.inputHeader)
		if !reflect.DeepEqual(tc.wantString, got) || !reflect.DeepEqual(tc.wantError, err) {
			t.Fatalf("expected: %v;%v, got: %v;%v", tc.wantString, tc.wantError, got, err)
		}

	}
}
