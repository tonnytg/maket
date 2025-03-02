package target

import "errors"

type Service struct {
	repo RepositoryInterface
}

type ServiceInterface interface {
	FindAll() ([]*Target, error)
	FindByUUID(uuid string) (*Target, error)
	Save(target *Target) error
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindAll() ([]*Target, error) {
	return s.repo.GetAll()
}

func (s *Service) FindByUUID(uuid string) (*Target, error) {
	return s.repo.GetByUUID(uuid)
}

func (s *Service) Save(target *Target) error {

	if _, err := target.Validate(); err != nil {
		return errors.New(ErrInvalidEntity)
	}

	if target.GetUUID() != "" {
		return s.repo.Update(target)
	}

	return s.repo.Create(target)
}
