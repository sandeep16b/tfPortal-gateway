package routes

import (
	"fmt"
	"net/http"
	"tfPortal-gateway/middleware"
	"tfPortal-gateway/proxy"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	fmt.Println("setup routes")

	// Public login route
	mux.HandleFunc("/login", middleware.LoginHandler)
	fmt.Println("completed handler file")
	// Protected proxy routes
	mux.Handle("/user/", middleware.JWTMiddleware(proxy.UserService()))
	mux.Handle("/post/", middleware.JWTMiddleware(proxy.PostService()))

	return mux
}
