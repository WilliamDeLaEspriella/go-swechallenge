package repository

import model "github.com/WilliamDeLaEspriella/go-swechallenge/models"

type RatingChangeRepositoryInterface interface {
	SelectRatingChange(query model.QueryRatingChange) []model.RatingChange
	InsertRatingChange(post model.PostRatingChange) bool
	SelectBestRatingChange() []model.RatingChange
}
