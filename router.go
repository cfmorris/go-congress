package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cfmorris/go-congress/config"
	"github.com/julienschmidt/httprouter"
)

func serveRoutes() {
	config, _ := config.ReadConfig("./config.yaml")

	router := httprouter.New()
	fmt.Printf("Server running on localhost: %s\n", config.Server.Port)
	router.GET("/", Index)
	router.GET("/senate", Index)
	router.GET("/house", Index)
	router.GET("/senate/:num", GetHouse)
	router.GET("/house/:num", GetHouse)
	router.GET("/bill/:num", Index)

	log.Fatal(http.ListenAndServe(":"+config.Server.Port, router))
}
