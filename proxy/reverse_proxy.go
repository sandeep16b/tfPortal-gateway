package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func UserService() http.Handler {
	fmt.Println("entered into userservice")
	url, _ := url.Parse("http://localhost:8093") // user_service
	return httputil.NewSingleHostReverseProxy(url)
}

func PostService() http.Handler {
	url, _ := url.Parse("http://localhost:8092") // post_service
	return httputil.NewSingleHostReverseProxy(url)
}
