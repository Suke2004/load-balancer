package reverseproxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(address string) *httputil.ReverseProxy {
	backendURL, err := url.Parse(address)

	if err !=nil{
		fmt.Println("URL not available")
	}

	return httputil.NewSingleHostReverseProxy(backendURL)
}

