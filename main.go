package main

import (
	"time"

	"googleSignOn/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	ict := time.Now().Local().Location()
	time.Local = ict
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routers.Setup(router) 
}