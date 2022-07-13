package middlewares

import (
	"net/http"
	"time"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func verifyGoogleAuthetication(idToken string) (*oauth2.Tokeninfo, bool) {

	httpClient := &http.Client{Timeout: 5 * time.Second}

	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, false
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, false
	}

	return tokenInfo, true
}