package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	model "github.com/WilliamDeLaEspriella/go-swechallenge/models"
	"github.com/WilliamDeLaEspriella/go-swechallenge/queries"
)

type RatingChangeRepository struct {
	DB *sql.DB
}

var validOrderColumns = map[string]bool{
	"created_at":  true,
	"ticker":      true,
	"target_to":   true,
	"target_from": true,
}

func NewRatingChangeRepository(db *sql.DB) RatingChangeRepositoryInterface {
	return &RatingChangeRepository{DB: db}
}

func (repository *RatingChangeRepository) InsertRatingChange(post model.PostRatingChange) bool {
	stmt, err := repository.DB.Prepare(queries.InsertRatingChange)
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(
		post.Ticker,
		post.Company,
		post.Brokerage,
		post.Action,
		post.RatingFrom,
		post.RatingTo,
		post.TargetFrom,
		post.TargetTo,
	)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func (repository *RatingChangeRepository) SelectRatingChange(query model.QueryRatingChange) []model.RatingChange {
	var result []model.RatingChange
	orderColumn := query.OrderBy // por ejemplo: "ticker"
	if !validOrderColumns[orderColumn] {
		orderColumn = "created_at" // fallback seguro
	}
	var (
		rows *sql.Rows
		err  error
	)
	if query.Search != "" {
		rows, err = repository.DB.Query(
			fmt.Sprintf(queries.GetRatingChangesBySearch,
				orderColumn,
				orderFormat(query.Order)),
			query.Search,
			query.Page,
			query.Offset)

	} else {
		rows, err = repository.DB.Query(
			fmt.Sprintf(queries.GetRatingChanges,
				orderColumn,
				orderFormat(query.Order)),
			query.Page,
			query.Offset)
	}

	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id         uint
			ticker     string
			company    string
			brokerage  string
			action     string
			ratingFrom string
			ratingTo   string
			targetFrom float64
			targetTo   float64
		)
		err := rows.Scan(
			&id,
			&ticker,
			&company,
			&brokerage,
			&action,
			&ratingFrom,
			&ratingTo,
			&targetFrom,
			&targetTo,
		)
		if err != nil {
			log.Println(err)
		} else {
			manga := model.RatingChange{
				Id:         id,
				Ticker:     ticker,
				Company:    company,
				Brokerage:  brokerage,
				Action:     action,
				RatingFrom: ratingFrom,
				RatingTo:   ratingTo,
				TargetFrom: targetFrom,
				TargetTo:   targetTo,
			}
			result = append(result, manga)
		}
	}
	return result
}

func orderFormat(order string) string {
	if strings.ToUpper(order) == "ASC" {
		return "ASC"
	}
	return "DESC"
}
func (repository *RatingChangeRepository) SelectBestRatingChange() []model.RatingChange {
	var result []model.RatingChange
	rows, err := repository.DB.Query(queries.BestRatingChange)
	if err != nil {
		log.Fatal("failed to execute query", err)
		return nil
	}
	for rows.Next() {
		var (
			id      uint
			ticker  string
			company string
		)
		err := rows.Scan(&ticker, &company)
		if err != nil {
			log.Println(err)
		} else {
			manga := model.RatingChange{
				Id:      id,
				Ticker:  ticker,
				Company: company,
			}
			result = append(result, manga)
		}
	}
	return result
}
