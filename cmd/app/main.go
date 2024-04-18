package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/lincolngondin/points-of-interest/internal/poi"
)

func main(){

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    poiHandler := poi.NewHandler()

    mux := http.NewServeMux()
    mux.HandleFunc("POST /pois", poiHandler.RegisterHandler)
    mux.HandleFunc("GET /pois", poiHandler.GetHandler)
    
    server := http.Server{
        Addr: ":3000",
        Handler: mux,
    }
    serverOpened := make(chan bool)


    go func(server *http.Server, serverOpened chan<-bool) {
        log.Println("Listening in port 3000!")
        err := server.ListenAndServe()
        if err == http.ErrServerClosed {
            log.Println("Server closed!")
        }
        serverOpened<-true
    }(&server, serverOpened)


    <-c
    server.Shutdown(context.Background())
    <-serverOpened

}
