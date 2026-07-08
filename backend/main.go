package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func main() {

	var err error

	db, err = pgx.Connect(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	_, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS visitors(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	)
	`)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/health", Health)
		api.GET("/visitors", GetVisitors)
		api.POST("/visitors", AddVisitor)
		api.DELETE("/visitors", DeleteVisitors)
	}

	log.Println("Server running on :8080")

	router.Run(":8080")
}
