package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oopss", http.StatusBadRequest)
			return
		}
		log.Printf("log: Data %s\n", d)
		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good bye Furki!")
	})

	http.ListenAndServe(":9000", nil)
}
