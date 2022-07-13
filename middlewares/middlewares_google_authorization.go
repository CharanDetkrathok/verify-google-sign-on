package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	errors_handler "googleSignOn/errors_handler"

	"github.com/gin-gonic/gin"
)

func GoogleAuthorization(c *gin.Context) {

		var googlePayload *GoogleAuthorizationPayload
		err := c.ShouldBindJSON(&googlePayload)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err)
			c.Abort()
			return 
		}

		oauth2Tokeninfo, ok := verifyGoogleAuthetication(googlePayload.IdToken)
		if !ok {
			c.IndentedJSON(http.StatusUnauthorized, errors_handler.NewVerifyGoogleUnauthorizedError())
			c.Abort()
			return
		}

		ok = strings.HasSuffix(oauth2Tokeninfo.Email, "@rumail.ru.ac.th")
		if !ok {
			c.IndentedJSON(http.StatusUnauthorized, errors_handler.NewNotRumailVerifyGoogleError())
			c.Abort()
			return
		}

		stdCode := strings.TrimSuffix(oauth2Tokeninfo.Email, "@rumail.ru.ac.th")
		if _, err = strconv.Atoi(stdCode); err != nil && len(stdCode) != 10 {
			c.IndentedJSON(http.StatusUnauthorized, errors_handler.NewNotStdCodeVerifyGoogleError())
			c.Abort()
			return 
		}

		googleTokeninfo := GoogleTokeninfo{
			StdCode: stdCode,
			Email: oauth2Tokeninfo.Email,
			VerifiedEmail: oauth2Tokeninfo.VerifiedEmail,
		}

		c.IndentedJSON(http.StatusOK, &googleTokeninfo)

}