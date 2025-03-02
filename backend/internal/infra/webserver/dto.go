package webserver

import "github.com/tonnytg/makemoneytarget/internal/domain/member"

type TargetDTO struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Status        string          `json:"status"`
	TargetAmount  float64         `json:"target_amount"`
	CurrentAmount float64         `json:"current_amount"`
	StartDate     string          `json:"start_date"`
	EndDate       string          `json:"end_date"`
	Members       []member.Member `json:"members"`
}
