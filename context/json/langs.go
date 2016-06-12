package main

import (
	"encoding/json"
	"fmt"
)

type Lang struct {
	LangName    string
	LangVersion string
}

type Langslice struct {
	Langs []Lang
}

func main() {
	var l Langslice
	str := `{"langs":[{"langName":"Golang", "langVersion":"2.0"},{"langName":"python", "langVersion":"3.5"}]}`
	json.Unmarshal([]byte(str), &l)
	fmt.Println(l)

	// json's data format is unkown
	b := []byte(`{"Name":"golang","Father":"google"}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("json error:", err)
	}
	fmt.Println(f)
	// type of f now is map[string]interface{}

	// to get data from f
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type unkown")
		}
	}

	// struct to JSON
	var j Langslice
	j.Langs = append(j.Langs, Lang{LangName: "C", LangVersion: "2.0"})
	j.Langs = append(j.Langs, Lang{LangName: "C++", LangVersion: "10.0"})
	k, err := json.Marshal(j)
	if err != nil {
		fmt.Println("json error: ", err)
	}
	fmt.Println(string(k))
}
