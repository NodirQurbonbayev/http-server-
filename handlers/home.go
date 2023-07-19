package handlers

import (
	"fmt"
	"net/http"
)

func Gethome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to Website server on your computer with your email address and password toconnect to the website")
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
