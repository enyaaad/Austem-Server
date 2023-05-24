package migrations

import (
	"AustemServer/models"
	"AustemServer/posgtresql/db"
	"fmt"
)

func init() {
	db.StartDB()
}

func Migrator() {
	err := db.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Project{})
	if err != nil {
		return
	}
	fmt.Println("migrations complete")
}
