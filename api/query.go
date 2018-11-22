package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/techfort/dis"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"golang.org/x/sync/errgroup"
)

func fuzzy(r *redis.Client, pattern string) (map[string][]string, error) {
	var cursor uint64
	str, sets, hashes := []string{}, []string{}, []string{}
	for {
		var keys []string
		var err error
		keys, cursor, err = r.Scan(cursor, fmt.Sprintf(`%v*`, pattern), 10).Result()
		fmt.Println(fmt.Sprintf(`k: %+v, c: %v`, keys, cursor))
		if err != nil {
			fmt.Println(err.Error())
		}
		g, _ := errgroup.WithContext(context.Background())

		for _, k := range keys {
			k := k
			g.Go(func() error {
				k := k
				s, err := r.Type(k).Result()
				switch s {
				case "string":
					str = append(str, k)
				case "set":
					sets = append(sets, k)
				case "hash":
					hashes = append(hashes, k)
				}
				return err
			})
		}
		err = g.Wait()
		if err != nil {
			fmt.Println(err.Error())
		}
		if cursor == 0 {
			break
		}
	}
	return map[string][]string{
		"string": str,
		"hash":   hashes,
		"set":    sets,
	}, nil

}

// Fuzzy performs a fuzzy search for keys
func Fuzzy(c echo.Context) error {
	cc := c.(*Context)
	pattern := c.Param("pattern")
	res, err := fuzzy(cc.Redis, pattern)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	json, err := json.Marshal(res)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	return cc.JSONBlob(http.StatusOK, []byte(json))
}

// Key returns a key from redis performing the correct retrieval operation
func Key(c echo.Context) error {
	cc := c.(*Context)
	key, typ := c.Param("key"), c.Param("type")
	var raw strResult
	switch typ {
	case "SET":
		str, err := dis.NewSet(cc.Redis, key).GetStrings()
		if err != nil {
			return cc.JSONBlob(http.StatusInternalServerError, Err(err))
		}
		raw = newStrResult(key, strings.Join(str, ", "), typ)
	case "HASH":
		str, err := dis.NewHash(cc.Redis, key).String()
		if err != nil {
			return cc.JSONBlob(http.StatusInternalServerError, Err(err))
		}
		raw = newStrResult(key, str, typ)
	case "STRING":
		str, err := dis.NewCache(cc.Redis).String(key)
		if err != nil {
			return cc.JSONBlob(http.StatusInternalServerError, Err(err))
		}
		raw = newStrResult(key, str, typ)
	}
	json, err := json.Marshal(raw)
	if err != nil {
		return cc.JSONBlob(http.StatusInternalServerError, Err(err))
	}
	return cc.JSONBlob(http.StatusOK, json)
}
