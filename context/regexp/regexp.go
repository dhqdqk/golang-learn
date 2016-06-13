package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); m {
		return true
	}
	return false
}

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// lower HTML TAG
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// delete style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// delete script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// delete html code
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// delete unless \n
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))

	a := "I am learning Golang"
	re, _ = regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个结果
	one := re.Find([]byte(a))
	fmt.Println("Find the one:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("Findall:", all)
	for _, v := range all {
		fmt.Println(string(v))
	}

	//查找符合条件的index位置，开始位置和结束位置;只匹配首次
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex:", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)
}
