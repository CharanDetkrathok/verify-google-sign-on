package middlewares

import "context"

var ctx = context.Background()

type (
	cors struct{}

	GoogleAuthorizationPayload struct {
		IdToken string `json:"id_token"`
	}

	GoogleTokeninfo struct {
		StdCode       string `json:"std_code,omitempty"`
		Email         string `json:"email,omitempty"`
		VerifiedEmail bool `json:"verified_email,omitempty"`
	}
)