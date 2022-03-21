package database

import (
	"service-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetSellers(db *gorm.DB) ([]models.Seller, error) {
	buyer := []models.Seller{}
	query := db.Select("sellers.*")
	if err := query.Find(&buyer).Error; err != nil {
		return buyer, err
	}
	return buyer, nil
}

func GetSellerByID(id string, db *gorm.DB) (models.Seller, bool, error) {
	b := models.Seller{}
	query := db.Select("sellers.*")
	err := query.Where("sellers.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteSeller(id string, db *gorm.DB) error {
	var b models.Seller
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSeller(db *gorm.DB, b *models.Seller) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
