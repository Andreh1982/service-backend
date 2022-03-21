package database

import (
	"service-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetBuyers(db *gorm.DB) ([]models.Buyer, error) {
	buyer := []models.Buyer{}
	query := db.Select("buyers.*")
	if err := query.Find(&buyer).Error; err != nil {
		return buyer, err
	}
	return buyer, nil
}

func GetBuyerByID(id string, db *gorm.DB) (models.Buyer, bool, error) {
	b := models.Buyer{}
	query := db.Select("buyers.*")
	err := query.Where("buyers.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteBuyer(id string, db *gorm.DB) error {
	var b models.Buyer
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBuyer(db *gorm.DB, b *models.Buyer) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
