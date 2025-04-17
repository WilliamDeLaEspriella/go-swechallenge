package finance

import (
	"context"
	"log"

	"github.com/WilliamDeLaEspriella/go-swechallenge/config"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

type Finance struct {
	Ticker string
}

func NewFinance(ticker string) FinanceInterface {
	return &Finance{Ticker: ticker}
}

func (finance *Finance) GetFinanceStock() *models.GetTickerDetailsResponse {

	c := polygon.New(config.Envs.POLYGON_API_KEY)
	params := models.GetTickerDetailsParams{
		Ticker: finance.Ticker,
	}

	res, err := c.GetTickerDetails(context.Background(), &params)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
