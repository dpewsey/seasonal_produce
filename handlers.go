// handlers.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"seasonal_produce/models"
)

func GetProduce(c *gin.Context) {
	var produce []models.Produce
	if err := DB.Find(&produce).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve produce"})
		return
	}

	c.JSON(http.StatusOK, produce)
}

// Implement other CRUD handlers (Create, Update, Delete) similarly
