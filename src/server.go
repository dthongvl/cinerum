package src

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return router
}

func StartWebServer(port string) {
	log.Println("Start at: " + port)

	http.Handle("/", NewRouter())
	http.HandleFunc("/ws", ServeWebSocket)

	http.ListenAndServe(":"+ port, nil)

	go handleMessages()
}