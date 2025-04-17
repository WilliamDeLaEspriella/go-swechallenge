package controller

import (
	"database/sql"
	"strconv"

	"github.com/WilliamDeLaEspriella/go-swechallenge/finance"
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
	pageStr := g.DefaultQuery("page", "1")
	limitStr := g.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	db := controller.DB
	repo_rating := repository.NewRatingChangeRepository(db)
	ratings_changes := repo_rating.SelectRatingChange(model.QueryRatingChange{
		Page:    limit,
		Offset:  offset,
		Search:  g.DefaultQuery("search", ""),
		Order:   g.DefaultQuery("order", "DESC"),
		OrderBy: g.DefaultQuery("orderBy", "created_at"),
	})
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

func (controller *RatingChangesController) BestRatingChanges(g *gin.Context) {
	db := controller.DB
	repo_rating := repository.NewRatingChangeRepository(db)
	ratings_changes := repo_rating.SelectBestRatingChange()
	if ratings_changes != nil {
		g.JSON(200, gin.H{"status": "success", "data": ratings_changes, "msg": "get ratings_changes successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get ratings_changes successfully"})
	}
}

func (controller *RatingChangesController) RatingChangesDetails(g *gin.Context) {
	id := g.Param("id")
	financeApi := finance.NewFinance(id)
	financeStock := financeApi.GetFinanceStock()
	if financeStock != nil {
		g.JSON(200, gin.H{"status": "success", "data": financeStock})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil})
	}
}
