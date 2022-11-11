package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	REQ_URL = "http://scanme.nmap.org/"
)

func main() {
	proxyFile, err := os.Open("http_proxies.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer proxyFile.Close()

	scanner := bufio.NewScanner(proxyFile)

	for scanner.Scan() {
		urlStr := scanner.Text()
		urlStr = "http://" + urlStr
		fmt.Println(urlStr)
		fixedURL, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		tr := &http.Transport{
			Proxy: http.ProxyURL(fixedURL),
		}
		client := http.Client{
			Transport: tr,
			Timeout:   time.Second * 10,
		}
		req, err := http.NewRequest("GET", REQ_URL, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(req.URL.String())
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("failed to get response")
		} else {
			fmt.Println("Got response from ", urlStr)
			defer resp.Body.Close()
		}
	}
}
