package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	// "strings"
)

func main() {
	// var url string
	fmt.Print("Введите URL: ")
	// fmt.Fscan(os.Stdin, &url)
	url := "http://116.203.203.76:3000/tasks/Циклическая ротация"
	HttpGetRequest(url)
}

func HttpGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n%s %s", url, body)
	// fmt.Printf("Count for \"%s\": %d \n", url, strings.Count(string(body), "Go"))
}
