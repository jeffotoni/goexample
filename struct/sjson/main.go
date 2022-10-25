package main

import (
	"encoding/json"
	"strconv"
)

type MyStructZero struct {
	ID          string `json:"ID,omitempty"`
	Description string `json:"description,omitempty"`
	Code        string `json:"Code,omitempty"`
}

type MyStructOne struct {
	ID            *int           `json:"ID"`
	Plan          string         `json:"Plan"`
	Loc           string         `json:"Loc"`
	Discord       *bool          `json:"Discord"`
	MyStructZero  *MyStructZero  `json:"MyStructZero"`
	MyStructZero2 []MyStructZero `json:"MyStructZero2,omitempty"`
}

type MyStructThree struct {
	ID      string `json:"ID,omitempty"`
	Plan    string `json:"Plan,omitempty"`
	Loc     string `json:"Loc,omitempty"`
	Discord *bool  `json:"Discord,omitempty"`

	MyStructZero  *MyStructZero  `json:"MyStructZero,omitempty"`
	MyStructZero2 []MyStructZero `json:"MyStructZero2,omitempty"`
}

func main() {

	jstr := `{"ID":0,"Plan":"xxxxx","Loc":"", "MyStructZero":{"Description":"test map json", "Code":"x388383"},"MyStructZero2":[{"ID":"xx99991"}]}`
	JsonM(jstr)
	println("...............................")
	jstr = `{"Plan":"xxxxx","Loc":"","Discord":null}`
	JsonM(jstr)
}

func JsonM(jstr string) {
	var a MyStructOne
	err := json.Unmarshal([]byte(jstr), &a)
	if err != nil {
		println("error:", err.Error())
		return
	}

	var cs MyStructThree
	if a.ID != nil {
		cs.ID = strconv.Itoa(*a.ID)
	}

	cs.Plan = a.Plan
	cs.Loc = a.Loc
	if a.Discord != nil {
		cs.Discord = a.Discord
	}
	if a.MyStructZero != nil {
		cs.MyStructZero = a.MyStructZero
	}

	for _, v := range a.MyStructZero2 {
		var of MyStructZero
		of.ID = v.ID
		cs.MyStructZero2 = append(cs.MyStructZero2, of)
	}

	b, _ := json.Marshal(&cs)
	println(string(b))

}
