package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, r.Body)
		r.Body.Close()
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}
