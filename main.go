package main

import (
	"flag"
	"net/http"

	"github.com/lucaspari/go-api/routes"
	"github.com/lucaspari/go-api/types"
)

func main(){
    port := flag.String("p",":3000","port for listening")
    flag.Parse()
    router := http.NewServeMux()
    routes.InitializeRoutes(router)
    server := types.NewServer(*port)
    server.Run(router)
}
