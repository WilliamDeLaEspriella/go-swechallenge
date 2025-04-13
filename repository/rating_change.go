package repository

import (
	"database/sql"
	"log"

	model "github.com/WilliamDeLaEspriella/go-swechallenge/models"
)

type RatingChangeRepository struct {
	DB *sql.DB
}

func NewRatingChangeRepository(db *sql.DB) RatingChangeRepositoryInterface {
	return &RatingChangeRepository{DB: db}
}

func (repository *RatingChangeRepository) InsertRatingChange(post model.PostRatingChange) bool {
	stmt, err := repository.DB.Prepare(
		`INSERT INTO rating_changes (
		    ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to
		) VALUES (
		    $1, $2, $3, $4, $5, $6, $7, $8
		)`,
	)
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

func (repository *RatingChangeRepository) SelectRatingChange() []model.RatingChange {
	var result []model.RatingChange
	rows, err := repository.DB.Query("SELECT * FROM rating_changes")
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
			targetFrom string
			targetTo   string
			createdAt  string
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
			&createdAt,
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
				CreatedAt:  createdAt,
			}
			result = append(result, manga)
		}
	}
	return result
}
