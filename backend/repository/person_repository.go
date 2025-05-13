// Método adicional para el PersonRepository
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