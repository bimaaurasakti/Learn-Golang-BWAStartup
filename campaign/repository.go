package campaign

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Campaign, error)
	GetByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) GetByUserID(userID int) ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Where("user_id = ?", userID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}