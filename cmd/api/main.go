package main

import (
	"github.com/spf13/viper"
	"github.com/techfort/dis"
	"github.com/techfort/dis/api"
)

/*
func main() {
	panic("not implemented")
}
*/

func main() {
	v := viper.New()
	v.AutomaticEnv()

	client, err := dis.NewRedisClient(v)

	if err != nil {
		panic(err)
	}
	_, err = api.InitAPI(client.Client)
	if err != nil {
		return
	}
}
