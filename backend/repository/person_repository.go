package repository

import (
	"errors"

	"github.com/Sepas8/death-note-app/backend/models"
	"gorm.io/gorm"
)

// PeopleRepository define el repositorio para Person
type PeopleRepository struct {
	db *gorm.DB
}

// NewPeopleRepository crea una nueva instancia del repositorio
func NewPeopleRepository(db *gorm.DB) *PeopleRepository {
	return &PeopleRepository{db: db}
}

// FindByFullName busca una persona por su nombre completo
func (p *PeopleRepository) FindByFullName(fullName string) (*models.Person, error) {
	var person models.Person
	err := p.db.Where("full_name = ?", fullName).First(&person).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}

// Create crea una nueva persona
func (p *PeopleRepository) Create(person *models.Person) error {
	return p.db.Create(person).Error
}

// FindByID busca una persona por su ID
func (p *PeopleRepository) FindByID(id uint) (*models.Person, error) {
	var person models.Person
	err := p.db.First(&person, id).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (p *PeopleRepository) FindAll() ([]*models.Person, error) {
	var people []*models.Person
	err := p.db.Find(&people).Error
	return people, err
}
