// db.go
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"seasonal_produce/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=seasonal_produce_app password=password dbname=seasonal_produce port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

	// Auto-migrate the schema
	db.AutoMigrate(&models.Produce{})
}
// CreateProduce adds a new produce item to the database
func CreateProduce(produce *models.Produce) {
	DB.Create(produce)
}

// GetProduce retrieves all produce items from the database
func GetProduce() []models.Produce {
	var produce []models.Produce
	DB.Find(&produce)
	return produce
}

// GetProduceByID retrieves a produce item by ID from the database
func GetProduceByID(id uint) (models.Produce, error) {
	var produce models.Produce
	result := DB.First(&produce, id)
	return produce, result.Error
}

// UpdateProduce updates a produce item in the database
func UpdateProduce(produce *models.Produce) {
	DB.Save(produce)
}

// DeleteProduce deletes a produce item from the database
func DeleteProduce(id uint) {
	DB.Delete(&models.Produce{}, id)
}
