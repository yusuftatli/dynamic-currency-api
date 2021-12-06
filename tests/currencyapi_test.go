package tests

import (
	"log"
	"testing"

	"github.com/yusuftatli/hepsiburada/config"
	"github.com/yusuftatli/hepsiburada/currencyapiclient"
	"github.com/yusuftatli/hepsiburada/httpclient"
	"github.com/yusuftatli/hepsiburada/models"
	cache "github.com/yusuftatli/hepsiburada/redis"
)


func TestSetGet(t *testing.T) {
	cfg, err := config.NewDefaultConfig()
	if err != nil {
		log.Panicln("failed to load config")
	}

	httpClient := httpclient.NewClient()
	currencyProvider := currencyapiclient.NewCurrencyProviderHandler(cfg, httpClient)
	currencyProvider.GetCurrenies()
	resp := &models.CurrencyModel{}
	cache.Initialize().GetKey("EURO", resp)  

	if resp.Currency == "EURO" && resp.Data == 11.4 {
		t.Errorf("value set incorrect")
	}
}