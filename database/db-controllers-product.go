package database

import (
	"service-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetProducts(db *gorm.DB) ([]models.Product, error) {
	buyer := []models.Product{}
	query := db.Select("products.*")
	if err := query.Find(&buyer).Error; err != nil {
		return buyer, err
	}
	return buyer, nil
}

func GetProductByID(id string, db *gorm.DB) (models.Product, bool, error) {
	b := models.Product{}
	query := db.Select("products.*")
	err := query.Where("products.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteProduct(id string, db *gorm.DB) error {
	var b models.Product
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, b *models.Product) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
