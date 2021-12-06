package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusuftatli/hepsiburada/config"
	"github.com/yusuftatli/hepsiburada/service"
)

func InitializeRoutes(env *config.Environment) {

	api := mux.NewRouter().PathPrefix("/").Subrouter()
	api.HandleFunc("/currency/{currencyCode}", service.GetCurrency).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", api))
}
