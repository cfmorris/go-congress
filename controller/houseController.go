package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func House(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	switch ps.ByName("house") {
	case "senate":
		fmt.Fprintf(w, "you've reached the %s", ps.ByName("house"))
	case "house":
		fmt.Fprintf(w, "you've reached the %s", ps.ByName("house"))
	default:
		fmt.Fprintf(w, "%s cannot be found", ps.ByName("house"))
	}
}
