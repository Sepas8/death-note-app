package repository

type Entity interface{}

type Repository[T Entity] interface {
	FindAll() ([]*T, error)
	FindById(id int) (*T, error)
	Save(*T) (*T, error)
	Delete(*T) error
}
