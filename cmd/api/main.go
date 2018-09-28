package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	apiHandlers "github.com/alexmourapb/base64api/cmd/api/handlers"
	apiRouters "github.com/alexmourapb/base64api/cmd/api/routers"
)

func main() {
	flag.Parse()

	fmt.Print(`
888888b.                               .d8888b.      d8888         d8888          d8b 
888  "88b                             d88P  Y88b    d8P888        d88888          Y8P 
888  .88P                             888          d8P 888       d88P888              
8888888K.   8888b.  .d8888b   .d88b.  888d888b.   d8P  888      d88P 888 88888b.  888 
888  "Y88b     "88b 88K      d8P  Y8b 888P "Y88b d88   888     d88P  888 888 "88b 888 
888    888 .d888888 "Y8888b. 88888888 888    888 8888888888   d88P   888 888  888 888 
888   d88P 888  888      X88 Y8b.     Y88b  d88P       888   d8888888888 888 d88P 888 
8888888P"  "Y888888  88888P'  "Y8888   "Y8888P"        888  d88P     888 88888P"  888 
                                                                         888          
                                                                         888          
                                                                         888          
`)
	port := "8181"

	fmt.Println("Porta:", port)
	h := &apiHandlers.Handler{}
	routes := apiRouters.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routes,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("can't initialize Server ->", err)
		os.Exit(1)
	}
}
