package category

import "github.com/google/uuid"

type createCategoryInput struct {
	Name string `json:"name" validate:"required,max=32" example:"electronics"`
}

type createBulkCategoryInput struct {
	Categories []string `json:"categories" validate:"required,dive,max=32" example:"electronics,furniture,school"`
}

type updateCategoryInput struct {
	OldName string `json:"-" validate:"required,max=32"`
	NewName string `json:"new_name" validate:"required,max=32" example:"electronics"`
}

type deleteCategoryInput struct {
	Name string `json:"-" validate:"required,max=32"`
}

type categoryOutput struct {
	Id        uuid.UUID `json:"id" example:"6ba7b814-9dad-11d1-80b4-00c04fd430c8"`
	Name      string    `json:"category" example:"electronics"`
	CreatedAt int64     `json:"created_at" example:"1709885754000"`
	UpdatedAt int64     `json:"updated_at" example:"1709885754000"`
}
