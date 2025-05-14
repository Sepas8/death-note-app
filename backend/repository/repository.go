package repository

import (
	"github.com/Sepas8/death-note-app/backend/models"
	"gorm.io/gorm"
)

type Entity interface {
	models.Person | models.Kill
}

type Repository[T Entity] struct {
	db *gorm.DB
}

func NewRepository[T Entity](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) FindAll() ([]*T, error) {
	var entities []*T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *Repository[T]) FindById(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Save(entity *T) (*T, error) {
	err := r.db.Save(entity).Error
	return entity, err
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}
