package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oopss", http.StatusBadRequest)
		return
	}
	log.Printf("log: Data %s\n", d)
	fmt.Fprintf(w, "Hello %s\n", d)
}
