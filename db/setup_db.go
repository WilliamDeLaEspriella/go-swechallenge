package db

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/WilliamDeLaEspriella/go-swechallenge/config"
	model "github.com/WilliamDeLaEspriella/go-swechallenge/models"
	"github.com/WilliamDeLaEspriella/go-swechallenge/repository"
)

type SetupDb struct {
	DB *sql.DB
}

func NewSetupDb(db *sql.DB) SetupDbInterface {
	return &SetupDb{DB: db}
}

func (setup *SetupDb) BulkRatingChanges() {
	nextPage := ""
	for {
		response := getExternalRatingChanges(nextPage)
		rawItems, ok := response["items"]
		if !ok {
			log.Fatal("Error setup db")
		}
		itemsJSON, err := json.Marshal(rawItems)
		if err != nil {
			log.Fatal("Error setup db")
		}
		var items []model.RatingChange
		if err := json.Unmarshal(itemsJSON, &items); err != nil {
			log.Fatal("Error setup db")
		}
		InsertExternalRatingChanges(items, setup.DB)
		np, ok := response["next_page"].(string)
		log.Println("Next page:", np)
		if !ok {
			log.Println("❌ 'next_page' no es string o no existe en la respuesta")
			break
		}

		if np == "" {
			break
		}

		nextPage = np

	}
}

func getExternalRatingChanges(next_page string) map[string]any {
	req, err := http.NewRequest("GET", config.Envs.SETUP_DB_URL+"?next_page="+next_page, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+config.Envs.SETUP_DB_TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func InsertExternalRatingChanges(items []model.RatingChange, db *sql.DB) {
	for _, r := range items {
		repo := repository.NewRatingChangeRepository(db)
		repo.InsertRatingChange(model.PostRatingChange{
			Ticker:     r.Ticker,
			Company:    r.Company,
			Brokerage:  r.Brokerage,
			Action:     r.Action,
			RatingFrom: r.RatingFrom,
			RatingTo:   r.RatingTo,
			TargetFrom: r.TargetFrom,
			TargetTo:   r.TargetTo,
		})
		log.Println("Company", r.Company, "inserted!!")
	}
}
