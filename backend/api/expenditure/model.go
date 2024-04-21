package expenditure

import "github.com/google/uuid"

type createExpenditureInput struct {
	Paisa      int32     `json:"paisa" validate:"required,gt=0" example:"561000"`
	CategoryId uuid.UUID `json:"category" validate:"required" example:"6ba7b814-9dad-11d1-80b4-00c04fd430c8"`
	UnixMillis int64     `json:"millis" validate:"gte=0" example:"1709885754000"`
}

type updateExpenditureInput struct {
	Id         string    `json:"-" validate:"required,uuid"`
	Paisa      int32     `json:"paisa" example:"561000"`
	CategoryId uuid.UUID `json:"category" validate:"required" example:"6ba7b814-9dad-11d1-80b4-00c04fd430c8"`
	UnixMillis int64     `json:"millis" validate:"gte=0" example:"1709885754000"`
}

type deleteExpenditureInput struct {
	Id string `json:"-" validate:"required,uuid"`
}

type getExpendituresInput struct {
	Page string `json:"-" validate:"required,int"`
}

type expenditureOutput struct {
	Id        uuid.UUID `json:"id" example:"6ba7b814-9dad-11d1-80b4-00c04fd430c8"`
	Paisa     int32     `json:"paisa" example:"123400"`
	Category  string    `json:"category" example:"electronics"`
	CreatedAt int64     `json:"created_at" example:"1709885754000"`
	UpdatedAt int64     `json:"updated_at" example:"1709885754000"`
}
