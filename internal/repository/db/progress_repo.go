package db

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/jinzhu/gorm"
)

type ProgressRepo struct {
	DB *gorm.DB
}

func NewProgressRepo(db *gorm.DB) *ProgressRepo {
	return &ProgressRepo{DB: db}
}

func (r *ProgressRepo) UpdateProgress(progress *models.Progress) error {
	return r.DB.Save(progress).Error
}

func (r *ProgressRepo) GetProgress(userID uint, language string) (*models.Progress, error) {
	var progress models.Progress

	err := r.DB.Where("user_id = ? AND language = ?", userID, language).First(&progress).Error
	return &progress, err
}
