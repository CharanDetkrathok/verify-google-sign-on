package routers

import (
	"net/http"

	"googleSignOn/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	router.Use(middlewares.NewCorsAccessControll().CorsAccessControll())

	router.GET("/healthz", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "The services work normally.",
		})
	})

	router.POST("/google-authorization", middlewares.GoogleAuthorization)

	router.Run(":8885")

}
