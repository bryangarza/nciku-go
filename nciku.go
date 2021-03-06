package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const searchurl string = "http://www.nciku.com/search/all/%v"
const swfurl string = "http://images.nciku.com/stroke_order/%v.swf"

func GetPage(c string) []byte {
	url := fmt.Sprintf(searchurl, c)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "Mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func SearchForID(c string) string {
	search := fmt.Sprintf("(\\d+)\">%v", c)
	r, err := regexp.Compile(search)
	if err != nil {
		log.Fatal(err)
	}
	page := GetPage(c)
	match := r.Find(page)
	// taking off the last 5 chars, ie ">好 is 5 chars long
	// Chinese characters take up 3 bytes
	match = match[:(len(match) - 5)]
	return string(match)
}

func StrokeURL(c string) string {
	URL := fmt.Sprintf(swfurl, SearchForID(c))
	return URL
}

func main() {
	// GetPage("好")
	fmt.Println(StrokeURL("好"))
}
