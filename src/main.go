package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		fmt.Println("url parse error", err)
		panic(err)
	}
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
