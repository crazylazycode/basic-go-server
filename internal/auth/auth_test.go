package auth

import (
	"net/http"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "secret"
	password2 := "supersecret"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := map[string]struct {
		password      string
		hash          string
		wantErr       bool
		matchPassword bool
	}{
		"Correct Password": {
			password:      password1,
			hash:          hash1,
			wantErr:       false,
			matchPassword: true,
		},
		"Incorrect Password": {
			password:      "wrongpassword",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		"Password doesn't match different hash": {
			password:      password1,
			hash:          hash2,
			wantErr:       false,
			matchPassword: false,
		},
		"Empty Password": {
			password:      "",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		"Invalid hash": {
			password:      password1,
			hash:          "invalidhash",
			wantErr:       true,
			matchPassword: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			match, err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && match != tt.matchPassword {
				t.Errorf("CheckPasswordHash() expects %v, got %v", tt.matchPassword, match)
			}
		})
	}

}

func TestGetBearerToken(t *testing.T) {

	tests := map[string]struct {
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		"Valid Bearer Token": {
			headers: http.Header{
				"Authorization": []string{"Bearer secret-token-123"},
			},
			wantToken: "secret-token-123",
			wantErr:   false,
		},
		"Missing Auhtorization Header": {
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		"Empty Authorization Header": {
			headers: http.Header{
				"Authorization": []string{""},
			},
			wantToken: "",
			wantErr:   true,
		},
		"Missing Bearer Prefix": {
			headers: http.Header{
				"Authorization": []string{"secret-token-123"},
			},
			wantToken: "",
			wantErr:   true,
		},
		"Wrong Auth scheme": {
			headers: http.Header{
				"Authorization": []string{"Apikey secret-token-123"},
			},
			wantToken: "",
			wantErr:   true,
		},
		"Bearer wihtout token": {
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && token != tt.wantToken {
				t.Errorf("GetBearerToken() expects %q, got %q", tt.wantToken, token)
			}
		})
	}
}
