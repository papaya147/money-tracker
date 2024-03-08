package expenditure

import "github.com/google/uuid"

type createExpenditureInput struct {
	Paisa    int32  `json:"paisa" validate:"required" example:"561000"`
	Category string `json:"category" validate:"required,max=32" example:"electronics"`
}

type expenditureOutput struct {
	Id        uuid.UUID `json:"id" example:"6ba7b814-9dad-11d1-80b4-00c04fd430c8"`
	Paisa     int32     `json:"paisa" example:"123400"`
	Category  string    `json:"category" example:"electronics"`
	CreatedAt int64     `json:"created_at" example:"1709885754000"`
	UpdatedAt int64     `json:"updated_at" example:"1709885754000"`
}
