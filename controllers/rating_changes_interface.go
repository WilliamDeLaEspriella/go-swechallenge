package controller

import "github.com/gin-gonic/gin"

type RatingChangesControllerInterface interface {
	InsertRatingChanges(g *gin.Context)
	GetRatingChanges(g *gin.Context)
	BestRatingChanges(g *gin.Context)
	RatingChangesDetails(g *gin.Context)
}
