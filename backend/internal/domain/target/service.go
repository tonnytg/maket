package target

import (
	"errors"
	"log"
)

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

func (s *Service) Save(target *Target) (*Target, error) {

	if _, err := target.Validate(); err != nil {
		return nil, err
	}

	t, _ := s.repo.GetByUUID(target.UUID)
	if t != nil {
		log.Println("target already exists")

		err := s.repo.Update(target)
		if err != nil {
			return nil, errors.New("error updating target")
		}
	}

	tt, err := s.repo.Create(target)
	if err != nil {
		return nil, errors.New("error creating target")
	}

	return tt, nil
}
