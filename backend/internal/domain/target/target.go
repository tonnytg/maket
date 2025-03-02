package target

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tonnytg/makemoneytarget/internal/domain/member"
)

const (
	TargetStatusActive   = "active"
	TargetStatusInactive = "inactive"

	ErrorInvalidUUID                 = "invalid uuid"
	ErrNameIsRequired                = "name is required"
	ErrDescriptionIsRequired         = "description is required"
	ErrTargetAmountIsRequired        = "target amount is required"
	ErrStartDateIsRequired           = "start date is required"
	ErrEndDateIsRequired             = "end date is required"
	ErrStartDateIsGreaterThanEndDate = "start date is greater than end date"

	ErrInvalidEntity = "invalid entity"
)

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

type TargetInterface interface {
	GetUUID() string
	GetName() string
	GetDescription() string
	GetStatus() string
	GetTargetAmount() float64
	GetCurrentAmount() float64
	GetMembers() []member.Member
	GetStartDate() string
	GetEndDate() string
	SetName(name string)
	SetDescription(description string)
	SetStatus(status string)
	SetEnable()
	SetDisable()
	SetTargetAmount(amount float64)
	SetNewCurrentAmount(amount float64)
	SetAppendedAmount(amount float64)
	SetStartDate(date string)
	SetEndDate(date string)
	SetAppendedMembers([]member.Member)
	SetNewMembers([]member.Member)
	Validate() bool
}

func (t Target) GetUUID() string {
	return t.UUID
}

func (t Target) GetName() string {
	return t.Name
}

func (t Target) GetDescription() string {
	return t.Description
}

func (t Target) GetTargetAmount() float64 {
	return t.TargetAmount
}

func (t Target) GetCurrentAmount() float64 {
	return t.CurrentAmount
}

func (t Target) GetMembers() []member.Member {
	return t.Members
}

func (t Target) GetStartDate() string {
	return t.StartDate
}

func (t Target) GetEndDate() string {
	return t.EndDate
}

func (t *Target) SetName(name string) {
	t.Name = name
}

func (t *Target) SetDescription(description string) {
	t.Description = description
}

func (t *Target) SetStatus(status string) {
	t.Status = status
}

func (t *Target) SetEnable() {
	t.SetStatus(TargetStatusActive)
}

func (t *Target) SetDisable() {
	t.SetStatus(TargetStatusInactive)
}

func (t *Target) SetTargetAmount(amount float64) {
	t.TargetAmount = amount
}

func (t *Target) SetNewCurrentAmount(amount float64) {
	t.CurrentAmount = amount
}

func (t *Target) SetAppendedAmount(amount float64) {
	t.CurrentAmount += amount
}

func (t *Target) SetStartDate(date string) {
	t.StartDate = date
}

func (t *Target) SetEndDate(date string) {
	t.EndDate = date
}

func (t *Target) SetAppendedMembers(members []member.Member) {
	t.Members = append(t.Members, members...)
}

func (t *Target) SetNewMembers(members []member.Member) {
	t.Members = members
}

func (t *Target) Validate() (*Target, error) {

	if t.Name == "" {
		return nil, fmt.Errorf(ErrNameIsRequired)
	}

	if t.Description == "" {
		return nil, fmt.Errorf(ErrDescriptionIsRequired)
	}

	if t.TargetAmount == 0 {
		return nil, fmt.Errorf(ErrTargetAmountIsRequired)
	}

	if t.StartDate == "" {
		return nil, fmt.Errorf(ErrStartDateIsRequired)
	}

	if t.EndDate == "" {
		return nil, fmt.Errorf(ErrEndDateIsRequired)
	}

	if t.StartDate > t.EndDate {
		return nil, fmt.Errorf(ErrStartDateIsGreaterThanEndDate)
	}

	return t, nil
}

func NewTarget(name, description, startDate, endDate string, targetAmount float64, members []member.Member) *Target {

	t := Target{}
	t.UUID = uuid.New().String()
	t.SetName(name)
	t.SetDescription(description)
	t.SetStatus(TargetStatusActive)
	t.SetNewCurrentAmount(0)
	t.SetStartDate(startDate)
	t.SetEndDate(endDate)
	t.SetTargetAmount(targetAmount)
	t.SetNewMembers(members)

	if _, err := t.Validate(); err != nil {
		return nil
	}

	return &t
}
