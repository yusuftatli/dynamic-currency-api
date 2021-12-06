package currencyapiclient

import (
	"encoding/json"
	"sort"
	"sync"

	"github.com/yusuftatli/hepsiburada/httpclient"
	"github.com/yusuftatli/hepsiburada/models"
	cache "github.com/yusuftatli/hepsiburada/redis"
)

type ICurrencyProvider interface {
	Handle(channel chan models.ProviderResponse, waitGroup *sync.WaitGroup)
}

type CurrencyProvider struct {
	cfg       models.ProviderConfig
	providers []ICurrencyProvider
}

//providers are registering
func NewCurrencyProviderHandler(cfg *models.ProviderConfig, httpClient *httpclient.Client) *CurrencyProvider {

	return &CurrencyProvider{
		cfg: *cfg,
		providers: []ICurrencyProvider{
			CurrencyProvider1{
				config:     cfg.Provider1,
				httpClient: *httpClient,
			},
			CurrencyProvider2{
				config:     cfg.Provider2,
				httpClient: *httpClient,
			},
			CurrencyProvider3{
				config:     cfg.Provider3,
				httpClient: *httpClient,
			},
		},
	}
}

func (currencyProvider *CurrencyProvider) GetCurrenies() {
	var waitGroup sync.WaitGroup
	channel := make(chan models.ProviderResponse)

	for _, provider := range currencyProvider.providers {
		waitGroup.Add(1)
		go provider.Handle(channel, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	var arrEur []float64
	var arrUsd []float64
	var arrGbp []float64
	//get results from channel and split by currency,  
	for msg := range channel {
		arrEur = append(arrEur, msg.Eur)
		arrUsd = append(arrUsd, msg.Usd)
		arrGbp = append(arrGbp, msg.Gbp)
	}
	eurModel := &models.CurrencyModel{Currency: "EURO", Data: sort.Float64Slice(arrEur)[0]}
	eurValue, _ := json.Marshal(eurModel)
	usdModel := &models.CurrencyModel{Currency: "USD", Data: sort.Float64Slice(arrUsd)[0]}
	usdValue, _ := json.Marshal(usdModel)
	gbpModel := &models.CurrencyModel{Currency: "GBP", Data: sort.Float64Slice(arrGbp)[0]}
	gbpValue, _ := json.Marshal(gbpModel)
	//write to redis
	cache.Initialize().GetKey("EURO", eurValue)
	cache.Initialize().GetKey("USD", usdValue)
	cache.Initialize().GetKey("GBP", gbpValue)
}
