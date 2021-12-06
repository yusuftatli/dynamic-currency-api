package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusuftatli/hepsiburada/models"
	cache "github.com/yusuftatli/hepsiburada/redis"
)

func GetCurrency(w http.ResponseWriter, r *http.Request) {
	currencyCode := mux.Vars(r)["currencyCode"]
	resp := &models.CurrencyModel{}
	cache.Initialize().GetKey(currencyCode, resp)
	models.RespondWithJSON(w, http.StatusOK, resp)
}
