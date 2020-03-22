package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// register a func to a path on http handler (defaultServeMux)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oop", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write(([]byte("Ooops")))
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// nil means using defaultServeMux
	http.ListenAndServe(":9090", nil)
}
