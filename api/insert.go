package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/techfort/dis"
)

// Set performs a SET
func Set(c echo.Context) error {
	key, value := c.Param("key"), c.Param("value")
	cc := c.(*Context)
	cache := dis.NewCache(cc.Redis)
	err := cache.StoreString(key, value)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	return cc.JSONBlob(http.StatusOK, []byte(OKMessage))
}

// HSet performs a HSET
func HSet(c echo.Context) error {
	key, field, value := c.Param("key"), c.Param("field"), c.Param("value")
	cc := c.(*Context)
	cache := dis.NewHash(cc.Redis, key)
	_, err := cache.StoreString(field, value)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	return cc.JSONBlob(http.StatusOK, []byte(OKMessage))
}

// SAdd performs a SADD
func SAdd(c echo.Context) error {
	key, value := c.Param("key"), c.Param("value")
	cc := c.(*Context)
	cache := dis.NewSet(cc.Redis, key)
	_, err := cache.AddString(value)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	return cc.JSONBlob(http.StatusOK, []byte(OKMessage))
}
