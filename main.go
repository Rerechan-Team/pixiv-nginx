package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, _ := url.Parse("http://127.0.0.1:80")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/websocket-http", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServeTLS(":443", "ca/pixiv.net.crt", "ca/pixiv.net.key", nil)
}