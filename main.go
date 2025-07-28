package golanghttprouter

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Hello Http Router!")
	})

	server := http.Server{
		Handler: router,
		Addr:    "Localhost:3000",
	}

	server.ListenAndServe()
}