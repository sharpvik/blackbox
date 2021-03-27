package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Made with ❤️  by sharpvik for Black Box")
	fmt.Println("My GitHub: https://github.com/sharpvik/")
	fmt.Println("BlackBox GitHub Repo: https://github.com/sharpvik/blackbox/")

	server := &http.Server{
		Addr:    ":8000",
		Handler: mainHandler(users),
	}

	fmt.Println("Serving at http://localhost:8000/ ...")
	server.ListenAndServe()
}
