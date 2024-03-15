# mxload_go

Golang middleware for Mx.Load
## Installation

```
go get github.com/mxload/mxload_go
```

## Usage

### go-chi

```golang
import (
    ...

    mxload "github.com/mxload/mxload_go"
)

func main() {
    r := chi.NewRouter()

    // Add mxload middleware
	r.Use(mxload.MiddlewareFunc(
        // mxload collector configuration
		&mxload.MxloadConfig{
			Enabled:      true, // Make collector enable/disable
			CollectorURL: "https://lambda-public.mxload.mx/put-request-log",
			ProjectID:    YOUR_PROJECT_ID, // Set your project ID
			ServiceKey:   "YOUR_SERVICE_KEY", // Set your service key
			CustomHeaders: []string{ // Set headers that should be sent to mxload.mx
				"Authorization",
			},
			IgnorePaths: []string{ // Set ignored paths
				"/ignored",
			},
		},
	))

    ....
}
```

### echo

```golang
import (
    ...

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
    ...
}
```

### gin example

```golang
import (
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	mxload "github.com/mxload/mxload_go"
)

func main() {
  r := gin.Default()

  r.Use(adapter.Wrap(mxload.MiddlewareFunc(
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
  ...
}
```
