package repository

type Entity interface {
	models.Person | models.Kill
}

type Repository[T Entity] interface {
	FindAll() ([]*T, error)
	FindById(id int) (*T, error)
	Save(*T) (*T, error)
	Delete(*T) error
}