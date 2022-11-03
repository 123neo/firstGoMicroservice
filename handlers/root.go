package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Root struct {
	l *log.Logger
}

func NewRoot(l *log.Logger) *Root {
	return &Root{l}
}

func (h *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Hello World")
	h.l.Println("Root logged")
	fmt.Fprintf(w, "Root here")
}
