// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://jeffpg:pass1234@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

}
