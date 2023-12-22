package main

import (
	"net/http/httputil"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}
