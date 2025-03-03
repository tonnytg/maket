package member

import "github.com/google/uuid"

type Member struct {
	UUID  string
	Name  string
	Email string
	Role  string
}

func NewMember(name, email, role string) *Member {

	return &Member{
		UUID:  uuid.New().String(),
		Name:  name,
		Email: email,
		Role:  role,
	}
}
