package services

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/chigaji/maron_speak/internal/repository/db"
)

type ProgressService struct {
	ProgressRepo *db.ProgressRepo
}

func NewProgressService(progressRepo *db.ProgressRepo) *ProgressService {
	return &ProgressService{ProgressRepo: progressRepo}
}

func (s *ProgressService) UpdateProgress(userID uint, language string, wordsLeaned uint) error {
	progress, err := s.ProgressRepo.GetProgress(userID, language)
	if err != nil {
		progress = &models.Progress{
			UserID:       userID,
			Language:     language,
			WordsLearned: wordsLeaned,
		}
	} else {
		progress.WordsLearned += wordsLeaned
	}
	return s.ProgressRepo.UpdateProgress(progress)
}

func (s *ProgressService) GetProgress(userID uint, language string) (*models.Progress, error) {
	return s.ProgressRepo.GetProgress(userID, language)
}
