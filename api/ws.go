package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/techfort/dis"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

type intResult struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type strResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func newIntResult(key string, i int64) intResult {
	return intResult{key, i}
}

func newStrResult(key, value, typ string) strResult {
	return strResult{key, value, typ}
}

// Ws is the websocket handler
func Ws(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	fmt.Println("starting ws...")
	cc := c.(*Context)
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	client := cc.Redis
	if err != nil {
		return err
	}
	defer ws.Close()
	ps := client.PSubscribe("__keyspace@0__:*")
	msgs := ps.Channel()
	cache := dis.NewCache(cc.Redis)
	for msg := range msgs {
		var raw interface{}
		var err error
		key := dis.Key(msg.Channel)
		op := msg.Payload
		fmt.Println(fmt.Sprintf("Message: %+v", msg))
		switch op {
		case "set":
			res, e := cache.String(key)
			if e != nil {
				err = e
			}
			raw = newStrResult(key, res, "KEY")
		case "hset":
			hash := dis.NewHash(cc.Redis, key)
			res, e := hash.String()
			if e != nil {
				err = e
			}
			raw = newStrResult(key, res, "HASH")
		case "sadd":
			set := dis.NewSet(cc.Redis, key)
			res, e := set.GetStrings()
			if err != nil {
				err = e
			}
			raw = newStrResult(key, strings.Join(res, ", "), "SET")
		default:
			fmt.Println("unknown operation")
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		bytes, err := json.Marshal(raw)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = ws.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}
	return nil
}
