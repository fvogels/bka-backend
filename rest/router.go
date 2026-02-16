package rest

import (
	"bass-backend/config"
	"bass-backend/database"
	"bass-backend/rest/routes/document"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

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
	router.Use(func(c *gin.Context) {
		log := slog.Default()

		log = log.With("start", time.Now())
		log = log.With("start", time.Now())
		log = log.With("path", c.Request.URL.Path)
		log = log.With("query", c.Request.URL.RawQuery)

		c.Next()

		status := c.Writer.Status()
		log = log.With("status", status)
		log = log.With("method", c.Request.Method)
		log = log.With("host", c.Request.Host)
		log = log.With("route", c.FullPath())
		log = log.With("end", time.Now())
		log = log.With("userAgent", c.Request.UserAgent())
		log = log.With("ip", c.ClientIP())
		log = log.With("referer", c.Request.Referer())

		isError := http.StatusBadRequest <= status
		if isError {
			log.Error("An error occurred while a request was handled")
		} else {
			log.Info("Request successfully handled")
		}
	})

	corsConfiguration := cors.DefaultConfig()
	corsConfiguration.AllowAllOrigins = true
	corsConfiguration.AllowCredentials = true
	router.Use(cors.New(corsConfiguration))

	return router
}

func (server *Server) defineEndPoints() {
	server.router.GET("/api/v1/documents", server.addDatabaseParameter(document.Handle))

	server.router.NoRoute(func(context *gin.Context) {
		context.File(config.HTMLPath)
	})
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
