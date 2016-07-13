package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	/*
		//fmt.Println(os.Args)
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
			os.Exit(1)
		}
	*/
	name := os.Args[0]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}
	os.Exit(0)
}
