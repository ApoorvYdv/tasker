package category

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// --- Create Category ---
type CreateCategoryPayload struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Color       *string `json:"color" validate:"omitempty,hexcolor"`
}

func (r *CreateCategoryPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Get Categories ---
type GetCategoriesQuery struct {
	Page   *int    `query:"page" validate:"omitempty,min=1"`
	Limit  *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Sort   *string `query:"sort" validate:"omitempty,oneof=created_at updated_at name"`
	Order  *string `query:"order" validate:"omitempty,oneof=asc desc"`
	Search *string `query:"search" validate:"omitempty,min=1"`
}

func (r *GetCategoriesQuery) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	// Set defaults for pagination
	if r.Page == nil {
		defaultPage := 1
		r.Page = &defaultPage
	}

	if r.Limit == nil {
		defaultLimit := 10
		r.Limit = &defaultLimit
	}

	if r.Sort == nil {
		defaultSort := "created_at"
		r.Sort = &defaultSort
	}

	if r.Order == nil {
		defaultOrder := "desc"
		r.Order = &defaultOrder
	}

	return nil
}

// --- Get Category by ID ---
type GetCategoryByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (r *GetCategoryByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Update Category ---
type UpdateCategoryPayload struct {
	ID          uuid.UUID `param:"id" validate:"required,uuid"`
	Name        *string   `json:"name" validate:"omitempty,min=3,max=100"`
	Description *string   `json:"description" validate:"omitempty,max=1000"`
	Color       *string   `json:"color" validate:"omitempty,hexcolor"`
}

func (r *UpdateCategoryPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Delete Category ---
type DeleteCategoryPayload struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (r *DeleteCategoryPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
