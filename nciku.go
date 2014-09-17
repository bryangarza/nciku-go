package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var mandarin = []string{
	"我",
	"叫",
	"什",
	"么",
	"名",
	"字",
}

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
	matchstr := string(match)
	return matchstr
}

func StrokeURL(mandarin []string) []string {
	ch := make(chan string)
	urls := make([]string, 1000)

	for _, c := range mandarin {
		go func(c string) {
			url := fmt.Sprintf(swfurl, SearchForID(c))
			ch <- url
		}(c)
	}

	for {
		select {
		case url := <-ch:
			fmt.Println(url)
			urls = append(urls, url)
			if len(urls) == len(mandarin) {
				return urls
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return urls
}

func main() {
	// GetPage("好")
	fmt.Println(StrokeURL(mandarin))
}
