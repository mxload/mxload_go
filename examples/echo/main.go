package main

import (
	"net/http"

	"github.com/labstack/echo"
	mxload "github.com/mxload/mxload_go"
)

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(mxload.MiddlewareFunc(
		&mxload.MxloadConfig{
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
