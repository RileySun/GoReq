package main

import(
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
	state string
	body string
}

func NewRequest() *Request {
	return &Request{
		response:&Response{
			state:"000",
			body:"",
		},
	}
}

func (r *Request) Send() (string, string, error) {
	//Check Method
	if r.method == "" {
		return "", "", errors.New("Missing Method")
	} 
	
	//Check URL
	if r.url == "" {
		return "", "", errors.New("Missing URL")
	}	
	if !isURL(r.url) {
		return "", "", errors.New("Invalid URL")
	}
	if r.url[0:8] != "https://" && r.url[0:7] != "http://" {
		if r.secure {
			r.url = "https://" + r.url
		} else {
			r.url = "http://" + r.url
		}
	}
	
	
	//Send Req
	status, response := httpReq(r.method, r.url, r.fields, r.headers)
	return status, response, nil
}