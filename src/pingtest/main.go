package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	var i int = 0
	ch := make(chan string)
	for i < 1000 {
		time.Sleep(time.Duration(time.Nanosecond * 1000000))
		go ping(ch)
		i++
	}

	for {
		fmt.Println(<-ch)
	}

	fmt.Println("post send success")
}

func ping(ch chan string) {
	//这里添加post的body内容
	data := make(map[string]string)
	data["workid"] = "1"
	data["machine"] = "2"

	str, _ := json.Marshal(data)

	resp, err := http.Post("http://localhost:9090/primary/apply",
		"application/json",
		strings.NewReader(string(str)))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	ch <- string(body)

	fmt.Println(string(body))
}
