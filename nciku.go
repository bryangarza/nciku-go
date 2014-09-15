package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func charid(c string) {
	resp, err := http.Get(fmt.Sprintf("http://www.nciku.com/search/all/%v", c))
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", string(body))
	}
}

func main() {
	charid("好")
}
