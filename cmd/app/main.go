package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/lincolngondin/points-of-interest/config"
	"github.com/lincolngondin/points-of-interest/internal/poi"
    _ "github.com/lib/pq"
)

func main() {
    configs := config.New()
    db, err := sql.Open(configs.DBDriverName, configs.DBDataSourceName)
    if err != nil {
        log.Fatal(err)
    }

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

    poiRepo := poi.NewRepository(db)
    poiService := poi.NewService(poiRepo)
	poiHandler := poi.NewHandler(poiService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /pois", poiHandler.RegisterHandler)
	mux.HandleFunc("GET /pois", poiHandler.GetHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	serverOpened := make(chan bool)

	go func(server *http.Server, serverOpened chan<- bool) {
		log.Println("Listening in port 3000!")
		err := server.ListenAndServe()
		if err == http.ErrServerClosed {
			log.Println("Server closed!")
		}
		serverOpened <- true
	}(&server, serverOpened)

	<-c
	server.Shutdown(context.Background())
	<-serverOpened

}
