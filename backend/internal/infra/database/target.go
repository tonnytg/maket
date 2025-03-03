package database

import (
	"github.com/tonnytg/makemoneytarget/internal/domain/member"
	"github.com/tonnytg/makemoneytarget/internal/domain/target"
	"log"
)

type TargetRepositorySqlite3 struct{}

func NewTargetRepositorySqlite3() *TargetRepositorySqlite3 {
	return &TargetRepositorySqlite3{}
}

type Target struct {
	UUID          string
	Name          string
	Description   string
	Status        string
	TargetAmount  float64
	CurrentAmount float64
	StartDate     string
	EndDate       string
	Members       []member.Member
}

func (r *TargetRepositorySqlite3) Create(t *target.Target) (*target.Target, error) {

	db := GetConnection()
	defer db.Close()

	// TODO: Melhor aqui
	newTarget := target.NewTarget(t.Name, t.Description, t.Status, t.StartDate, t.EndDate, t.TargetAmount, t.Members)

	log.Println("Creating target", newTarget)
	stmt, err := db.Prepare("INSERT INTO targets(uuid, name, description, status, target_amount, current_amount, start_date, end_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = stmt.Exec(newTarget.UUID, newTarget.Name, newTarget.Description, newTarget.Status, newTarget.TargetAmount, newTarget.CurrentAmount, newTarget.StartDate, newTarget.EndDate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Creating members")
	for i, m := range newTarget.Members {
		log.Println("Creating member: ", m)

		newMember := member.NewMember(m.Name, m.Email, m.Role)

		stmt, err = db.Prepare("INSERT INTO target_members(uuid, name, email, role, target_id) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		_, err = stmt.Exec(newMember.UUID, newMember.Name, newMember.Email, newMember.Role, newTarget.UUID)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		newTarget.Members[i] = *newMember
	}

	return newTarget, nil
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
