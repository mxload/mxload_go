package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	buuurst_dev "git.drecom.jp/diet/buuurst_dev_go"
)

func main() {
  r := gin.Default()

  r.Use(adapter.Wrap(buuurst_dev.MiddlewareFunc(
	&buuurst_dev.BuuurstDevConfig{
		Enabled:      true,
		CollectorURL: "https://lambda-public.buuurst.dev/put-request-log",
		ProjectID:    YOUR_PROJECT_ID,
		ServiceKey:   "YOUR_SERVICE_KEY",
		CustomHeaders: []string{
			"Authorization",
		},
		IgnorePaths: []string{
			"/ignored",
		},
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
