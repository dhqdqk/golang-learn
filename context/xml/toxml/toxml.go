package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Langs struct {
	XMLName xml.Name `xml:"langs"`
	Version string   `xml:"version,attr"`
	Svs     []lang   `xml:"lang"`
}

type lang struct {
	LangName    string `xml:"langName"`
	LangVersion string `xml:"langVersion"`
}

func main() {
	v := &Langs{Version: "1"}
	v.Svs = append(v.Svs, lang{"c", "1.0"})
	v.Svs = append(v.Svs, lang{"C++", "3.0"})
	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
