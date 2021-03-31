package main

import (
	"log"
	"net/http"
)

func main() {
	greet()

	server := &http.Server{
		Addr:    ":8000",
		Handler: mainHandler(users),
	}

	log.Print("Serving at http://localhost:8000/ ...")
	log.Print("Go to http://localhost:8000/api/Viktor to make sure it works!")
	server.ListenAndServe()
}

func greet() {
	log.Print("Made with ❤️  by sharpvik for BlackBox")
	log.Print("My GitHub: https://github.com/sharpvik/")
	log.Print("BlackBox GitHub Repo: https://github.com/sharpvik/blackbox/")
}
