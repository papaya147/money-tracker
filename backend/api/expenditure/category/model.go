package category

type createCategoryInput struct {
	Name string `json:"name" validate:"required,max=32" example:"electronics"`
}

type categoryOutput struct {
	Name      string `json:"category" example:"electronics"`
	CreatedAt int64  `json:"created_at" example:"1709885754000"`
	UpdatedAt int64  `json:"updated_at" example:"1709885754000"`
}
