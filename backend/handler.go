package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Visitor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateVisitor struct {
	Name string `json:"name" binding:"required"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func AddVisitor(c *gin.Context) {

	var req CreateVisitor

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var id int

	err := db.QueryRow(
		context.Background(),
		`INSERT INTO visitors(name)
		 VALUES($1)
		 RETURNING id`,
		req.Name,
	).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"name":    req.Name,
		"message": "visitor created",
	})
}

func GetVisitors(c *gin.Context) {

	rows, err := db.Query(
		context.Background(),
		`SELECT id, name
		 FROM visitors
		 ORDER BY id`,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	// Jangan gunakan nil slice
	visitors := make([]Visitor, 0)

	for rows.Next() {

		var visitor Visitor

		if err := rows.Scan(&visitor.ID, &visitor.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		visitors = append(visitors, visitor)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, visitors)
}

func DeleteVisitors(c *gin.Context) {

	_, err := db.Exec(
		context.Background(),
		"DELETE FROM visitors",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all visitors deleted",
	})
}
