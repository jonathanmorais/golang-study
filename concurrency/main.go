package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	go request("http://google.com/")
	request("http://google.com/")
	go request("http://bing.com/")
	request("http://bing.com/")
}

func request(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	t1 := time.Now()
	tf := time.Since(t1)

	time.Sleep(100 * time.Millisecond)
	log.Println(resp.StatusCode)
	log.Println(tf)
}
