package finance

import "github.com/polygon-io/client-go/rest/models"

type FinanceInterface interface {
	GetFinanceStock() *models.GetTickerDetailsResponse
}
