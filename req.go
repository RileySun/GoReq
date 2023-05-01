package main

import(
	"log"
	"errors"
)

type Request struct {
	method string
	url string
	headers []string
	fields []string
	secure bool
	
	response *Response
}

type Response struct {
	status string
	body string
}

func NewRequest() *Request {
	return &Request{
		headers:[]string{"Content-Type=application/x-www-form-urlencoded"},
		response:&Response{
			status:"000",
			body:"",
		},
	}
}

func (r *Request) Send() error {
	//Check Method
	if r.method == "" {
		return errors.New("Missing Method")
	} 
	
	//Check URL
	if r.url == "" {
		return errors.New("Missing URL")
	}	
	if !isURL(r.url) {
		return errors.New("Invalid URL")
	}
	if len(r.url) < 7 || (r.url[0:8] != "https://" && r.url[0:7] != "http://") {
		log.Println("Missing http schema")
		if r.secure {
			r.url = "https://" + r.url
		} else {
			r.url = "http://" + r.url
		}
	}
	
	//Send Req
	r.response.status, r.response.body = httpReq(r.method, r.url, r.fields, r.headers)
	return nil
}