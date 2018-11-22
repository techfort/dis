package dis

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type r struct {
	*redis.Client
}

type set struct {
	r
	key string
}

type hash struct {
	r
	key string
}

type disclient struct {
	r
}

// NewCache creates a cache struct for plain key value pairs
func NewCache(rd *redis.Client) Cache {
	return cache{r{rd}}
}

// NewSet returns a new redis set
func NewSet(rd *redis.Client, key string) Set {
	r := r{rd}
	return set{r, key}
}

// NewHash returns a hash struct for hash values
func NewHash(rd *redis.Client, key string) Hash {
	r := r{rd}
	return hash{r, key}
}

// Wrap wraps errors.Wrap
func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

// ParseInt is a wrapper for int64
func ParseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// Key replaces the redis prefix so the plain key of a notification is returned
func Key(s string) string {
	return strings.Replace(s, "__keyspace@0__:", "", 1)
}
