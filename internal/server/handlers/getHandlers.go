package handlers

import "github.com/gorilla/mux"

func GetAllRoutets() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/orders", OrdersHandler)
	r.HandleFunc("/orders/{id:[0-9a-zA-Z]+}", GetByIdOrderHandler)

	return r
}
