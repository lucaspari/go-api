package routes

import "net/http"

func InitializeRoutes(mux *http.ServeMux){
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}
