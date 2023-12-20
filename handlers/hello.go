package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Creamos una estrutura que impementa la interface http handler

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Esta funcion hace que la estructura Hello que creamos pueda implementar la interfaz
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	// read the body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(w, "Hello %s\n", b)

}
