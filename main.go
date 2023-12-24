package main

import (
	"embed"
  "log"
  "encoding/json"
  "github.com/gin-gonic/gin"
	"seasonal_produce/routes"
)
//go:embed data.json
var content embed.FS

type Produce struct {
    ID      int    `json:"id" gorm:"primaryKey"`
    Name    string `json:"name"`
    Season  string `json:"season"`
    Group   string `json:"group"` // think citrus
    Specific_Group string `json:"specific_group"` // think apple/oranages
    Start_Month  string `json:"start_month"`
    End_Month    string `json:"end_month"`
    Type    string `json:type` // fruit or veg
}

func main() {
  filePath := "data.json"
  produceData, err := content.ReadFile(filePath)
  if err != nil {
    log.Fatal("Failed to read configuration file:", err)
  }

  var produceList []Produce 
  err = json.Unmarshal(produceData, &produceList)
  if err != nil { 
    log.Fatal("failed to unmarshal JSON data:", err)
  }
	// Set up the router
	router := routes.SetupRouter()
	// Example route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the seasonal produce app!",
		})
	})
  router.GET("/produce", func(c *gin.Context) {
    c.JSON(200, produceList)
  })
	// Run the server
	router.Run(":8080")
}
