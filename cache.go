package dis

type cache struct {
	r
}

// Cache interface for storing plain keys in redis
type Cache interface {
	String(key string) (string, error)
	Int(key string) (int64, error)
	StoreInt(key string, i int64) error
	StoreString(key string, val string) error
	StoreInts(key string, i ...int64) error
	Delete(key string) error
}

func (c cache) String(key string) (string, error) {
	return c.Get(key).Result()
}

func (c cache) Int(key string) (int64, error) {
	r, err := c.Get(key).Result()
	if err != nil {
		return 0, Wrap(err, "failed to retrieve value from redis")
	}
	i, err := ParseInt(r)
	return i, Wrap(err, "failed to convert string to int")
}

func (c cache) StoreInt(key string, i int64) error {
	_, err := c.Set(key, i, 0).Result()
	return Wrap(err, "failed to store int")
}

func (c cache) StoreString(key string, val string) error {
	_, err := c.Set(key, val, 0).Result()
	return Wrap(err, "failed to store string")
}

func (c cache) StoreInts(key string, i ...int64) error {
	_, err := c.Set(key, i, 0).Result()
	return Wrap(err, "failed to store ints")
}

func (c cache) Delete(key string) error {
	_, err := c.Del(key).Result()
	return Wrap(err, "failed to delete key")
}
