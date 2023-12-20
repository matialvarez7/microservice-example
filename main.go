package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matialvarez7/microservice-example/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// referencia al handler products
	ph := handlers.NewProducts(l)

	// Creamos un nuevo server mux que registra los distintos handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Creamos un servidor al cual le configuramos las características que deseamos para nuestro servicio
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Encerramos el bloque en una goroutine para que no bloquee
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recevied terminate, graceful shutdown", sig)
	// Espera a que las request que ya estén hechas se completen antes de apagar el servidor
	timeOutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeOutContext)
}
