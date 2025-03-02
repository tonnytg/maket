package member

import "github.com/google/uuid"

type Member struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Role     string
	CreateAt string
	UpdateAt string
	DeleteAt string
}
