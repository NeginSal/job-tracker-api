package job

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(job *Job) error {
	return r.db.Create(job).Error
}

func (r *Repository) GetAllByUser(userID string) ([]Job, error) {
	var jobs []Job
	err := r.db.Where("user_id = ?", userID).Find(&jobs).Error
	return jobs, err
}

func (r *Repository) Delete(id uuid.UUID, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Job{}).Error
}
