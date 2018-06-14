package main

import "errors"

type Named interface {
	Name() string
}

func greeting(thing Named) (string, error) {
	if thing == nil {
		return "", errors.New("thing cannot be nil")
	}

	return "Hello " + thing.Name(), nil
}

/* error func
func greeting(thing Named) string {
	return "Hello " + thing.Name()
} */

func main() {
	greeting(nil)
}
