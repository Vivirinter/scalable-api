package main

import (
	"github.com/Vivirinter/scalable-api/internal/config"
	"github.com/Vivirinter/scalable-api/internal/db"
	"github.com/Vivirinter/scalable-api/internal/handler"
	"github.com/Vivirinter/scalable-api/internal/migrations"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	err = db.InitDB(cfg)
	if err != nil {
		return err
	}
	defer db.CloseDB()

	migrations.RunMigrations()

	r := gin.Default()
	r.GET("/board", handler.GetBoard())
	r.POST("/board", handler.PostBoard())

	log.Println("running...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return r.Run(":" + port)
}
