package api

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/techfort/dis"
)

const (
	// ErrMessage errors message
	ErrMessage = `{ "error": "%v"}`
	// OKMessage ok message
	OKMessage = `{ "status": "OK" }`
)

// Err wrapper for jsonblob with error
func Err(err error) []byte {
	return []byte(fmt.Sprintf(ErrMessage, err.Error()))
}

// Context is the dis Context for API requests
type Context struct {
	echo.Context
	Redis *redis.Client
}

// RoutesGET returns get routes
func RoutesGET() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		"/ws":             Ws,
		"/query/:pattern": Fuzzy,
		"/key/:key/:type": Key,
	}
}

// RoutesPOST returns POST routes
func RoutesPOST() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		"/set/:key/:value":         Set,
		"/sadd/:key/:value":        SAdd,
		"/hset/:key/:field/:value": HSet,
	}
}

// InitAPI starts the API
func InitAPI(r *redis.Client) (*echo.Echo, error) {
	e := echo.New()
	e = Config(e, r)
	err := e.Start(":1666")
	return e, dis.Wrap(err, "failed to start API")
}

// Config provides middleware and wraps the context of each request
func Config(e *echo.Echo, r *redis.Client) *echo.Echo {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c, r}
			return h(cc)
		}
	})
	for route, handler := range RoutesGET() {
		e.GET(route, handler)
	}
	for route, handler := range RoutesPOST() {
		e.POST(route, handler)
	}
	return e
}
