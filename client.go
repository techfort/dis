package dis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// NewRedisClient returns a redis client
func NewRedisClient(v *viper.Viper) (r, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     v.GetString(EnvRedisURL),
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	return r{client}, err
}
