package target

type Repository struct{}

type RepositoryInterface interface {
	Create(target *Target) (*Target, error)
	GetAll() ([]*Target, error)
	GetByUUID(uuid string) (*Target, error)
	Update(target *Target) error
	Delete(uuid string) error
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(target *Target) error {
	return nil
}

func (r *Repository) GetAll() ([]*Target, error) {
	return nil, nil
}

func (r *Repository) GetByUUID(uuid string) (*Target, error) {
	return nil, nil
}

func (r *Repository) Update(target *Target) error {
	return nil
}

func (r *Repository) Delete(uuid string) error {
	return nil
}
