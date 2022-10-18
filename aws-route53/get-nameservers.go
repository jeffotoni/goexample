package main

import (
	"fmt"
	"net"
	"strings"
)

func GetNS(d string) ([]string, error) {
	var nss []string
	n, _ := net.LookupNS(d)

	for _, v := range n {
		a := strings.TrimSuffix(v.Host, ".")
		nss = append(nss, a)
	}

	return nss, nil
}

func main() {
	nss, err := GetNS("google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nss)
}
