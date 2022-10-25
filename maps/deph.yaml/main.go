package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

var LocaleNotFound = errors.New("i18n-not-found")

var file = []byte(`
teste: oiamigo
teste_nest:
  chave: valor
  chaveint: 12
`)

type LangDict map[interface{}]interface{}

func (t LangDict) T(fullKey string) (interface{}, error) {
	splitKeys := strings.Split(fullKey, ".")
	depth := len(splitKeys)

	curr := map[interface{}]interface{}(t)
	for level, key := range splitKeys {
		fmt.Printf("level: %d; depth: %d; curr: %+v;\n", level, depth, curr)
		if next, ok := curr[key]; ok {
			fmt.Printf("next: %+v\n", next)
			if level+1 == depth {
				return next, nil
			}

			if curr, ok = next.(LangDict); !ok {
				fmt.Printf("Conversion failed\n")
				fmt.Println(curr)
				return nil, LocaleNotFound
			}
		}
	}

	return nil, LocaleNotFound
}

func loadDict() (LangDict, error) {
	dict := make(LangDict)
	err := yaml.Unmarshal(file, &dict)
	if err != nil {
		return nil, err
	}

	return dict, nil
}

func main() {
	dict, _ := loadDict()
	// fmt.Printf("%+v\n", dict)
	// out1, _ := dict.T("teste")
	// fmt.Printf("dict.T(\"teste\") -> %v", out1)
	out2, _ := dict.T("teste_nest.chave")
	fmt.Printf("dict.T(\"teste_nest\") -> %v", out2)
}
