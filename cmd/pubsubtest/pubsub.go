package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/techfort/dis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	ps := client.PSubscribe("__key*__:test*")
	msgs := ps.Channel()
	for msg := range msgs {
		fmt.Println(dis.Key(msg.Channel), msg.Payload)
	}
}
