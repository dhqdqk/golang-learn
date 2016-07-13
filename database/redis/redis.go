package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	// datatype string
	client.Set("a", []byte("hello"))
	val, _ := client.Get("a")
	fmt.Println(string(val))
	client.Del("a")

	// list
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("l")
}
