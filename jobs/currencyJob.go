package jobs

import (
	"strconv"
	"time"

	"github.com/yusuftatli/hepsiburada/config"
	"github.com/yusuftatli/hepsiburada/currencyapiclient"
	"github.com/yusuftatli/hepsiburada/httpclient"
	"github.com/yusuftatli/hepsiburada/models"
)

//Calling requests with timer
func StartCurrencyJob(cfg *models.ProviderConfig) {
	httpClient := httpclient.NewClient()
	currencyProvider := currencyapiclient.NewCurrencyProviderHandler(cfg, httpClient)

	timeInterval, err := strconv.Atoi(config.GetEnvironment().TimeInterval)
	if err != nil {
		timeInterval = 10
	}
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	currencyProvider.GetCurrenies()
	go func() {
		for range ticker.C {
			currencyProvider.GetCurrenies()
		}
	}()
}
