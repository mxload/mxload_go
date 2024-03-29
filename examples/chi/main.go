package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mxload "github.com/mxload/mxload_go"
)

func main() {
	r := chi.NewRouter()
	r.Use(mxload.MiddlewareFunc(
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
	))
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get"))
	})
	r.Post("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post"))
	})
	http.ListenAndServe(":3000", r)
}
