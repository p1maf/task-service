package task

import (
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) Get(id uint32) (Task, error) {
	var t Task
	if err := r.db.First(&t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Task{}, nil
		}
		return Task{}, err
	}
	return t, nil
}

func (r *Repository) List() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) ListByUser(userID uint32) ([]Task, error) {
	var tasks []Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) Update(task Task) (Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) Delete(id uint32) error {
	return r.db.Delete(&Task{}, id).Error
}
