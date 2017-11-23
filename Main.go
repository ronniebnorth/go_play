package main

import (
	//"time"
	"net/http"
	"fmt"
)


func main() {
	var r int
	for c := 0; c < 10000; c++ {
		r += c
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
	/*
	start := time.Now()
	http.Get("https://blog.archive.org/2016/10/23/defining-web-pages-web-sites-and-web-captures/")

	fmt.Println(time.Since(start))
	*/
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I lovess %s!", r.URL.Path[1:])
}