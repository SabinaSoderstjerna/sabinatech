package main

import (
	"fmt"

	"github.com/SabinaSoderstjerna/sabinatech/internal/handler"
	"github.com/SabinaSoderstjerna/sabinatech/internal/website"
)

func main() {
	aboutHandler := website.InitAboutHandler()
	indexHandler := website.InitIndexHandler()
	httpMux := website.InitHTTPMux([]*handler.Handler{aboutHandler, indexHandler})
	httpServer := website.InitHTTPServer(httpMux)

	fmt.Println("Listening")
	fmt.Println(httpServer.ListenAndServe())
}
