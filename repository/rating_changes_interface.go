package repository

import model "github.com/WilliamDeLaEspriella/go-swechallenge/models"

type RatingChangeRepositoryInterface interface {
	SelectRatingChange() []model.RatingChange
	InsertRatingChange(post model.PostRatingChange) bool
}
