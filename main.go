package main

import (
	"fmt"
	"net/http"

	"github.com/1amkaizen/jwtweb/controler"
)

func main() {
	http.HandleFunc("/", controler.Index)
	http.HandleFunc("/index", controler.Index)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)

}
