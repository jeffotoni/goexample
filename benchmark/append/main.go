package main

var igetThis = []byte(`{"code":"","a":null}`)

type A struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type B struct {
	Code string `json:"code"`
	A    []A    `json:"a"`
}

func elBSlaice(b B) bool {
	if b.A == nil {
		b.A = []A{{}}
	}
	return true
}

func elBAppend(b B) bool {
	if b.A == nil {
		var a A
		b.A = append(b.A, a)
	}
	return true
}

func main() {
	var b B
	elBSlaice(b)
	elBAppend(b)
}
