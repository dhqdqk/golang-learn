package main

import (
	"fmt"
	"os"
)

func main() {
	// dir
	os.Mkdir("golang", 0777)
	os.MkdirAll("golang/test1", 0777)
	err := os.Remove("golang")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("golang")

	// file
	userFile := "golang.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}
