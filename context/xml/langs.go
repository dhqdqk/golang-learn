package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlylangs struct {
	XMLName     xml.Name `xml:"langs"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"lang"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName      xml.Name `xml:"lang"`
	LangName     string   `xml:"langName"`
	LangVerseion string   `xml:"langVersion"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}

func main() {
	file, err := os.Open("langs.xml")
	checkErr(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	checkErr(err)
	v := Recurlylangs{}
	err = xml.Unmarshal(data, &v)
	checkErr(err)
	fmt.Println(v)
}
