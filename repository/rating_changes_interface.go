package repository

import model "github.com/WilliamDeLaEspriella/go-swechallenge/models"

type RatingChangeRepositoryInterface interface {
	SelectRatingChange(limit int, offset int) []model.RatingChange
	InsertRatingChange(post model.PostRatingChange) bool
}
