package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/WilliamDeLaEspriella/go-swechallenge/config"
	controller "github.com/WilliamDeLaEspriella/go-swechallenge/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (server *Server) CreateConnection() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full",
		config.Envs.POSTGRES_USER,
		config.Envs.POSTGRES_PASSWORD,
		config.Envs.POSTGRES_URI,
		config.Envs.POSTGRES_DB,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
	server.DB = db
}

func (server *Server) Migrate() {
	if _, err := server.DB.Exec(
		"CREATE TABLE IF NOT EXISTS rating_changes (id SERIAL NOT NULL PRIMARY KEY,ticker VARCHAR(10) NOT NULL,company VARCHAR(100) NOT NULL,brokerage VARCHAR(100) NOT NULL,action VARCHAR(20) NOT NULL,rating_from VARCHAR(50),rating_to VARCHAR(50),target_from DECIMAL(10, 2),target_to DECIMAL(10, 2),created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"); err != nil {
		log.Fatal(err)
		log.Println("ERROR", err)
	}
}

func (server *Server) CreateRoutes() {
	gin.SetMode(config.Envs.GIN_MODE)
	routes := gin.Default()
	controller := controller.NewRatingChangesController(server.DB)
	routes.GET("/rating_changes", controller.GetRatingChanges)
	routes.POST("/rating_changes", controller.InsertRatingChanges)
	server.Routes = routes
}

func (server *Server) Run() {
	server.Routes.Run(":" + config.Envs.PORT)
}
