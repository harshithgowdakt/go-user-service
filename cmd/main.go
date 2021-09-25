package main

import (
	"io"
	"log"
	"net/http"
)

func main()  {
	helloHandler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "hellow, world")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening to request at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}