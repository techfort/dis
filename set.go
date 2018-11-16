package dis

import (
	"github.com/pkg/errors"
)

// Set is the redis set interface
type Set interface {
	Key() string
	AddInt(i ...int64) (int64, error)
	AddString(i ...string) (int64, error)
	GetInts() ([]int64, error)
	GetStrings() ([]string, error)
	RemoveInt(i int64) (int64, error)
	RemoveString(i string) (int64, error)
	IsMemberInt(i int64) (bool, error)
	IsMemberString(i string) (bool, error)
	Union(newkey string, i Set) (Set, error)
	Diff(newkey string, i Set) (Set, error)
	Intersect(newkey string, i Set) (Set, error)
}

func (s set) Key() string {
	return s.key
}

func (s set) AddInt(i ...int64) (int64, error) {
	r, err := s.SAdd(s.key, i).Result()
	return r, errors.Wrap(err, "failed to add int to set")
}

func (s set) AddString(i ...string) (int64, error) {
	r, err := s.SAdd(s.key, i).Result()
	return r, errors.Wrap(err, "failed to add string to set")
}

func (s set) GetInts() ([]int64, error) {
	r, err := s.SMembers(s.key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve ints from redis set")
	}
	result, err := StrToIntSlice(r)
	return result, errors.Wrap(err, "failed to convert strings to ints")
}

func (s set) GetStrings() ([]string, error) {
	r, err := s.SMembers(s.key).Result()
	return r, errors.Wrap(err, "failed to retrieve string from set")
}

func (s set) RemoveInt(i int64) (int64, error) {
	r, err := s.SRem(s.key, i).Result()
	return r, errors.Wrap(err, "failed to remove int from set")
}

func (s set) RemoveString(i string) (int64, error) {
	r, err := s.SRem(s.key, i).Result()
	return r, errors.Wrap(err, "failed to remove string from set")
}

func (s set) IsMemberInt(i int64) (bool, error) {
	r, err := s.SIsMember(s.key, i).Result()
	return r, errors.Wrap(err, "failed to check if int is member of set")
}

func (s set) IsMemberString(i string) (bool, error) {
	r, err := s.SIsMember(s.key, i).Result()
	return r, errors.Wrap(err, "failed to check if int is member of set")
}

func (s set) Union(newkey string, i Set) (Set, error) {
	_, err := s.SUnionStore(s.key, i.Key()).Result()
	return NewSet(s.Client, newkey), errors.Wrap(err, "failed to perform a union")
}

func (s set) Diff(newkey string, i Set) (Set, error) {
	_, err := s.SDiffStore(s.key, i.Key()).Result()
	return NewSet(s.Client, newkey), errors.Wrap(err, "failed to perform a set diff")
}

func (s set) Intersect(newkey string, i Set) (Set, error) {
	_, err := s.SInterStore(s.key, i.Key()).Result()
	return NewSet(s.Client, newkey), errors.Wrap(err, "failed to perform an interset")
}
