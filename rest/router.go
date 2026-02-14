package rest

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/rest/routes/documents"
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	database *sql.DB
	router   *gin.Engine
}

func StartServer() error {
	db, err := database.OpenDatabase(config.DatabasePath)
	if err != nil {
		return err
	}

	server := newServer(db)
	defer server.shutdown()

	if err := server.run(); err != nil {
		return err
	}

	return nil
}

func newServer(database *sql.DB) *Server {
	server := Server{
		database: database,
		router:   createGinRouter(),
	}

	server.defineEndPoints()

	return &server
}

func createGinRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	corsConfiguration := cors.DefaultConfig()
	corsConfiguration.AllowAllOrigins = true
	corsConfiguration.AllowCredentials = true
	router.Use(cors.New(corsConfiguration))

	return router
}

func (server *Server) defineEndPoints() {
	server.router.GET("/api/v1/documents", server.addDatabaseParameter(documents.Handle))
}

func (server *Server) run() error {
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := server.router.Run(address); err != nil {
		return err
	}

	return nil
}

func (server *Server) shutdown() {
	server.database.Close()
}

func (server *Server) addDatabaseParameter(f func(*sql.DB, *gin.Context)) func(*gin.Context) {
	return func(context *gin.Context) {
		f(server.database, context)
	}
}
