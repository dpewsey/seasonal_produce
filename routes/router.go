// routes.go
package routes

import (
    "github.com/gin-gonic/gin"
//	  "seasonal_produce/db"
//	  "seasonal_produce/models"
//    "strconv"
  )

func SetupRouter() *gin.Engine {
    router := gin.Default()
  
//    router.GET("/produce", GetProduce)
//  	router.GET("/produce/:id", GetProduceByID)
//  	router.POST("/produce", CreateProduce)
//  	router.PUT("/produce/:id", UpdateProduce)
//  	router.DELETE("/produce/:id", DeleteProduce)


    return router
}
//// GetProduce handles the GET /produce route
//func GetProduce(c *gin.Context) {
//	produce := db.GetProduce()
//	c.JSON(200, produce)
//}
//
//// GetProduceByID handles the GET /produce/:id route
//func GetProduceByID(c *gin.Context) {
//	id := c.Params.ByName("id")
//	produceID := ConvertStringToUint(id)
//
//	produce, err := db.GetProduceByID(produceID)
//	if err != nil {
//		c.JSON(404, gin.H{"error": "Produce not found"})
//		return
//	}
//
//	c.JSON(200, produce)
//}
//
//// CreateProduce handles the POST /produce route
//func CreateProduce(c *gin.Context) {
//	var produce models.Produce
//	if err := c.ShouldBindJSON(&produce); err != nil {
//		c.JSON(400, gin.H{"error": "Invalid input"})
//		return
//	}
//
//	db.CreateProduce(&produce)
//	c.JSON(201, produce)
//}
//
//// UpdateProduce handles the PUT /produce/:id route
//func UpdateProduce(c *gin.Context) {
//	id := c.Params.ByName("id")
//	produceID := ConvertStringToUint(id)
//
//	produce, err := db.GetProduceByID(produceID)
//	if err != nil {
//		c.JSON(404, gin.H{"error": "Produce not found"})
//		return
//	}
//
//	if err := c.ShouldBindJSON(&produce); err != nil {
//		c.JSON(400, gin.H{"error": "Invalid input"})
//		return
//	}
//
//	db.UpdateProduce(&produce)
//	c.JSON(200, produce)
//}
//
//// DeleteProduce handles the DELETE /produce/:id route
//func DeleteProduce(c *gin.Context) {
//	id := c.Params.ByName("id")
//	produceID := ConvertStringToUint(id)
//
//	db.DeleteProduce(produceID)
//	c.JSON(204, gin.H{})
//}
//
//// ConvertStringToUint converts a string to a uint
//func ConvertStringToUint(s string) uint {
//	// Convert the string to a uint, handle errors as needed
//	result, err := strconv.ParseUint(s, 10, 0)
//	if err != nil {
//		// Handle the error, for example, return 0 or another default value
//		return 0
//	}
//
//	return uint(result)
//}
