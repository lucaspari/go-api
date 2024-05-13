package main

import (
	"net/http"

	"github.com/lucaspari/go-api/routes"
	"github.com/lucaspari/go-api/types"
)

func main(){
    router := http.NewServeMux()
    routes.InitializeRoutes(router)
    server := types.NewServer(":3000")
    server.Run(router)
}
