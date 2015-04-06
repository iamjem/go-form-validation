package forms

import (
	"strconv"
	"time"
)

type LoaderFunc func(string) (interface{}, error)

var StringLoader = LoaderFunc(func(rawValue string) (interface{}, error) {
	return rawValue, nil
})

var IntLoader = LoaderFunc(func(rawValue string) (interface{}, error) {
	val, err := strconv.ParseInt(rawValue, 0, 0)
	if err != nil {
		return nil, err
	}
	return int(val), nil
})

var TimeLoader = NewTimeLoader(time.RFC3339)

func NewTimeLoader(layout string) LoaderFunc {
	return LoaderFunc(func(rawValue string) (interface{}, error) {
		val, err := time.Parse(layout, rawValue)
		if err != nil {
			return nil, err
		}
		return val, nil
	})
}
