package dis

import (
	"encoding/json"
)

// Hash interface
type Hash interface {
	Key() string
	String() (string, error)
	StoreInt(field string, i int64) (bool, error)
	StoreString(field, val string) (bool, error)
	Incr(field string) (int64, error)
	IncrBy(field string, i int64) (int64, error)
	Delete(field ...string) error
}

func (h hash) Key() string {
	return h.key
}

func (h hash) String() (string, error) {
	r, err := h.HGetAll(h.key).Result()
	if err != nil {
		return "", Wrap(err, "failed to retrieve hash")
	}
	json, err := json.Marshal(r)
	if err != nil {
		return "", Wrap(err, "failed to marshall hash")
	}
	return string(json), err

}

func (h hash) StoreInt(field string, i int64) (bool, error) {
	r, err := h.HSet(h.key, field, i).Result()
	return r, Wrap(err, "failed to store value in hash")
}

func (h hash) StoreString(field, val string) (bool, error) {
	r, err := h.HSet(h.key, field, val).Result()
	return r, Wrap(err, "failed to store value in hash")
}

func (h hash) Incr(field string) (int64, error) {
	r, err := h.HIncrBy(h.key, field, 1).Result()
	return r, Wrap(err, "failed to increment value in hash")
}

func (h hash) IncrBy(field string, i int64) (int64, error) {
	r, err := h.HIncrBy(h.key, field, i).Result()
	return r, Wrap(err, "failed to increment value in hash")
}

func (h hash) Delete(field ...string) error {
	_, err := h.HDel(h.key, field...).Result()
	return Wrap(err, "failed to delete keys from hash")
}
