package dis

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

// StrToIntSlice converts an arary of strings into an array of int64
func StrToIntSlice(s []string) ([]int64, error) {
	g, _ := errgroup.WithContext(context.Background())
	r := make([]int64, len(s))
	for i, str := range s {
		i, str := i, str
		g.Go(func() error {
			i, str := i, str
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			r[i] = val
			return err
		})
	}
	err := g.Wait()
	return r, err
}

// ResToInts wraps a redis result for direct return
func ResToInts(s []string, err error) ([]int64, error) {
	r, e := StrToIntSlice(s)
	return r, errors.Wrap(err, e.Error())
}

func wi(i int64, err error, msg string) (int64, error) {
	return i, errors.Wrap(err, msg)
}

func ws(i string, err error, msg string) (string, error) {
	return i, errors.Wrap(err, msg)
}

func wis(i []int64, err error, msg string) ([]int64, error) {
	return i, errors.Wrap(err, msg)
}

func wss(i []string, err error, msg string) ([]string, error) {
	return i, errors.Wrap(err, msg)
}

func wb(b bool, err error, msg string) (bool, error) {
	return b, errors.Wrap(err, msg)
}
