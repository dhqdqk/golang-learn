package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	// set port of redis
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	// string type
	client.Set("a", []byte("hello"))
	val, _ := client.Get("a")
	fmt.Println(string(val))
	client.Del("a")

	// list range
	vals := []string{"a", "b", "c", "e"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		fmt.Println(i, ":", string(v))
	}

	client.Del("l")
}
