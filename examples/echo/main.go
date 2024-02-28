package main

import (
	"net/http"

	buuurst_dev "git.drecom.jp/diet/buuurst_dev_go"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(buuurst_dev.MiddlewareFunc(
		&buuurst_dev.BuuurstDevConfig{
			Enabled:      true,
			CollectorURL: "https://lambda-public.mxload.mx/put-request-log",
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

	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "get")
	})
	e.POST("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "post")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
