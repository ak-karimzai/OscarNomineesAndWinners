package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
	"github.com/nmferoz/db/token"
	"github.com/nmferoz/db/util"
)

type Server struct {
	config     util.Config
	query      db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, database *sql.DB) (*Server, error) {
	tokenMaker := token.NewJWTMaker(config.TokenSymmetricKey)

	server := &Server{
		config:     config,
		query:      *db.New(database),
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/movies", server.CreateMovie)
	authRoutes.PUT("/movies", server.UpdateMovie)
	router.GET("/movies", server.ListMovies)
	router.GET("/movies/:id", server.GetMovie)

	authRoutes.POST("/actors", server.CreateActor)
	authRoutes.PUT("/actors", server.UpdateActor)
	router.GET("/actors", server.ListActors)
	router.GET("/actors/:id", server.GetActor)

	authRoutes.POST("/awards", server.CreateAward)
	authRoutes.PUT("/awards", server.UpdateAward)
	router.GET("/awards", server.ListAwards)
	router.GET("/awards/:id", server.GetAward)

	authRoutes.POST("/nominations", server.CreateNomination)
	authRoutes.PUT("/nominations", server.UpdateNomination)
	router.GET("/nominations", server.ListNominations)
	router.GET("/nominations/:id", server.GetNomination)

	authRoutes.POST("/performances", server.CreatePerformance)
	authRoutes.PUT("/performances", server.UpdatePerformance)
	router.GET("/performances", server.ListPerformances)
	router.GET("/performances/:id", server.GetPerformance)

	server.router = router
}

func (server *Server) Run(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
