package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	buuurst_dev "github.com/mxload/buuurst_dev_go"
)

func main() {
	r := gin.Default()

  r.Use(adapter.Wrap(buuurst_dev.MiddlewareFunc(
	&buuurst_dev.BuuurstDevConfig{
		Enabled:      true,
		CollectorURL: "https://lambda-public.mxload.mx/put-request-log",
		ProjectID:    YOUR_PROJECT_ID,
		ServiceKey:   "YOUR_SERVICE_KEY",
		CustomHeaders: []string{
			"Authorization",
		},
	)))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	r.Run()
}
