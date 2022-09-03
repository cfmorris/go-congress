package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cfmorris/go-congress/update"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func House(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	num := ps.ByName("num")
	switch strings.ToLower(num) {
	case "senate":
		fmt.Fprintf(w, "you've reached the %s", ps.ByName("house"))
		//start db interaction
	case "house":
		fmt.Fprintf(w, "you've reached the %s", ps.ByName("house"))
	default:
		fmt.Fprintf(w, "%s cannot be found", ps.ByName("house"))
	}
}

func GetHouse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s Number: %s\n", ps.ByName("house"), ps.ByName("num"))
}

func main() {

	update.CheckUpdate()

	serveRoutes()
}
