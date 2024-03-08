package category

type createCategoryInput struct {
	Name string `json:"name" validate:"required,max=15"`
}

type categoryOutput struct {
	Name      string `json:"category"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
