package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"computerextra/elaerning-go/db"
	"computerextra/elaerning-go/env"
	"computerextra/elaerning-go/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Static Assets
	var dir string
	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	// Get DB Client
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	routes.GetRoutes(router, client)
	routes.GetApiRoutes(router, client)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(router)

	env := env.GetEnv()

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%v", env.PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
