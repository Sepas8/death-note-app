package repository

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type PeopleRepository struct {
	db *gorm.DB
}

func NewPeopleRepository(db *gorm.DB) *PeopleRepository {
	return &PeopleRepository{
		db: db,
	}
}

func (p *PeopleRepository) FindAll() ([]*models.Person, error) {
	var people []*models.Person
	err := p.db.Find(&people).Error
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (p *PeopleRepository) Save(data *models.Person) (*models.Person, error) {
	err := p.db.Save(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PeopleRepository) FindById(id int) (*models.Person, error) {
	var person models.Person
	err := p.db.Where("id = ?", id).First(&person).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &person, nil
}

func (p *PeopleRepository) Delete(data *models.Person) error {
	err := p.db.Delete(data).Error
	if err != nil {
		return err
	}
	return nil
}
