package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// Document source as returned by Elasticsearch
	b := json.RawMessage(`{
            "a_b": "c",
            "d_e": ["d"],
            "e_f": {
                    "g_h": {
                            "i_j": "k",
                            "l_m": {}
                    }
            }
    }`)

	x := convertKeys(b)

	buf := &bytes.Buffer{}
	json.Indent(buf, []byte(x), "", "  ")
	fmt.Println(buf.String())
}

func convertKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		// Not a JSON object
		return j
	}

	for k, v := range m {
		fixed := fixKey(k)
		delete(m, k)
		m[fixed] = convertKeys(v)
	}

	b, err := json.Marshal(m)
	if err != nil {
		return j
	}

	return json.RawMessage(b)
}

func fixKey(key string) string {
	return strings.ToUpper(key)
}
