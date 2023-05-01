package main

import(
	"log"
	"time"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"net/http"
	"net/url"
)

func isURL(e string) bool {
    urlRegex := regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
    return urlRegex.MatchString(e)
}

func getHostname(newUrl string) string {
    url, err := url.Parse(newUrl)
    if err != nil {
        log.Fatal(err)
    }
    return strings.TrimPrefix(url.Hostname(), "www.")
}

func httpReq(method string, reqURL string, body []string, headers []string) (string, string) {
	//Generic GET request to Open Weather Map API
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	
	//fields
	data := url.Values{}
	for _, field := range body {
		kv := strings.Split(field, "=")
		//must be a key value pair
		if len(kv) == 2 {
			data.Add(kv[0], kv[1])
		}
	}
	
	//Body
	bytesObj := []byte(data.Encode())
	newBody := bytes.NewBuffer(bytesObj)
	
	req, err := http.NewRequest(method, reqURL, newBody)
	if err != nil {
		log.Fatal(err)
	}
	
	
	//headers
	for _, header := range headers {
		kv := strings.Split(header, "=")
		//must be a key value pair
		if len(kv) == 2 {
			req.Header.Add(kv[0], kv[1])
		}
	}
	
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	responseData, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	finalResponse := string(responseData[:])
	
	defer resp.Body.Close()
	
	return resp.Status, finalResponse
}