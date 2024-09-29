package handler

import (
	"github.com/Vivirinter/scalable-api/internal/db"
	"github.com/Vivirinter/scalable-api/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetScalable() ([]model.Scalable, error) {
	const q = `SELECT author, content, created_at FROM scalable ORDER BY created_at DESC LIMIT 100`

	rows, err := db.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]model.Scalable, 0)

	for rows.Next() {
		var author string
		var content string
		var createAt time.Time
		err = rows.Scan(&author, &content, &createAt)
		if err != nil {
			return nil, err
		}
		results = append(results, model.Scalable{author, content, createAt})
	}

	return results, nil
}

func AddScalable(scalable model.Scalable) error {
	const q = `INSERT INTO scalable(author, content, created_at) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(q, scalable.Author, scalable.Content, scalable.CreatedAt)
	return err
}

func GetBoard() gin.HandlerFunc {
	return func(context *gin.Context) {
		results, err := GetScalable()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error " + err.Error()})
			return
		}
		context.JSON(http.StatusOK, results)
	}
}

func PostBoard() gin.HandlerFunc {
	return func(context *gin.Context) {
		var s model.Scalable
		if context.Bind(&s) == nil {
			s.CreatedAt = time.Now()
			if err := AddScalable(s); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error " + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	}
}
