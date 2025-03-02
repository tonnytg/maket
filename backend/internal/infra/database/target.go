package database

import (
	"github.com/tonnytg/makemoneytarget/internal/domain/target"
	"log"
)

type TargetRepositorySqlite3 struct{}

func NewTargetRepositorySqlite3() *TargetRepositorySqlite3 {
	return &TargetRepositorySqlite3{}
}

func (r *TargetRepositorySqlite3) Create(target *target.Target) (*target.Target, error) {
	log.Println("Creating target")
	return nil, nil
}

func (r *TargetRepositorySqlite3) GetAll() ([]*target.Target, error) {
	return nil, nil
}

func (r *TargetRepositorySqlite3) GetByUUID(uuid string) (*target.Target, error) {
	return nil, nil
}

func (r *TargetRepositorySqlite3) Update(target *target.Target) error {
	log.Println("Updating target")
	return nil
}

func (r *TargetRepositorySqlite3) Delete(uuid string) error {
	return nil
}
