package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("x returned", x())
	fmt.Println("y returned", y())
	fmt.Println("z returned", z())
}

func x() (err error) {
	{
		s, err := "", errors.New("something else 1")
		defer func() {
			err = errors.New("something")
		}()
		ignore(s)
		return err
	}
}

func y() (err error) {
	{
		s1, err := "", errors.New("something else 2")
		defer func() {
			err = errors.New("something")
		}()
		ignore(s1)
		return err
	}
}

func z() (err error) {
	s, err := "", errors.New("something else")
	{
		defer func() {
			err = errors.New("something")
		}()
		ignore(s)
		return err
	}
}

func ignore(s string) {}
