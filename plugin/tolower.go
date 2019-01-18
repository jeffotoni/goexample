// Go in action
// @jeffotoni
// 2019-01-18

package main

import (
	"strings"
)

type lower string

func (t lower) MustLower(name string) string {

	return strings.ToLower(name)
}

// exported as symbol "
var Tolower lower
