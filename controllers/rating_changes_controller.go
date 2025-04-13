package controller

import (
	"database/sql"
	"log"

	model "github.com/WilliamDeLaEspriella/go-swechallenge/models"
	"github.com/WilliamDeLaEspriella/go-swechallenge/repository"
	"github.com/gin-gonic/gin"
)

type RatingChangesController struct {
	DB *sql.DB
}

func NewRatingChangesController(db *sql.DB) RatingChangesControllerInterface {
	return &RatingChangesController{DB: db}
}

func (controller *RatingChangesController) GetRatingChanges(g *gin.Context) {
	db := controller.DB
	repo_rating := repository.NewRatingChangeRepository(db)
	ratings_changes := repo_rating.SelectRatingChange()
	log.Println("ratings_changes", ratings_changes)
	if ratings_changes != nil {
		g.JSON(200, gin.H{"status": "success", "data": ratings_changes, "msg": "get ratings_changes successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get ratings_changes successfully"})
	}
}

func (m *RatingChangesController) InsertRatingChanges(g *gin.Context) {
	db := m.DB
	var post model.PostRatingChange
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_rating := repository.NewRatingChangeRepository(db)
		insert := repo_rating.InsertRatingChange(post)
		if insert {
			g.JSON(200, gin.H{"status": "success", "msg": "insert ratings_changes successfully"})
		} else {
			g.JSON(500, gin.H{"status": "failed", "msg": "insert ratings_changes failed"})
		}
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}
