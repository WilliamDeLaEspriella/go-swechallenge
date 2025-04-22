package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/WilliamDeLaEspriella/go-swechallenge/config"
	controller "github.com/WilliamDeLaEspriella/go-swechallenge/controllers"
	"github.com/WilliamDeLaEspriella/go-swechallenge/db"
	"github.com/WilliamDeLaEspriella/go-swechallenge/queries"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	DB     *sql.DB
	Routes *gin.Engine
}

var ginLambda *ginadapter.GinLambdaV2

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

func (server *Server) CreateTables() {
	if _, err := server.DB.Exec(queries.CreateTables); err != nil {
		log.Fatal(err)
		log.Println("ERROR", err)
	}
}

func (server *Server) Migrate() {
	var count int
	err := server.DB.QueryRow(queries.CountRatingChange).Scan(&count)
	if err != nil {
		log.Fatal("failed to execute setup db query", err)
	}
	if count == 0 {
		setupDb := db.NewSetupDb(server.DB)
		setupDb.BulkRatingChanges()
	}
}

func (server *Server) ConfigCors() {
	gin.SetMode(config.Envs.GIN_MODE)
	ginGonic := gin.Default()
	ginGonic.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://vue2-swechallenge-9wnqaoo6t-william-de-la-espriellas-projects.vercel.app"}, // cambia según tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.Routes = ginGonic
}

func (server *Server) CreateRoutes() {
	gin.SetMode(config.Envs.GIN_MODE)
	routes := server.Routes
	controller := controller.NewRatingChangesController(server.DB)
	routes.GET("/rating_changes", controller.GetRatingChanges)
	routes.POST("/rating_changes", controller.InsertRatingChanges)
	routes.GET("/rating_changes/recommendation", controller.BestRatingChanges)
	routes.GET("/rating_changes/:id", controller.RatingChangesDetails)
	ginLambda = ginadapter.NewV2(server.Routes)

}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}
func (server *Server) Run() {
	//server.Routes.Run(":" + config.Envs.PORT)
	lambda.Start(Handler)
}
