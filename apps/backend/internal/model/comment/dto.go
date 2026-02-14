package comment

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// --- Create Comment ---
type CreateCommentRequest struct {
	TodoID  uuid.UUID `json:"todoId" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max=1000"`
}

func (r *CreateCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Get Comments ---
type GetCommentsRequest struct {
	TodoID uuid.UUID `param:"todoId" validate:"required,uuid"`
}

func (r *GetCommentsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Update Comment ---
type UpdateCommentRequest struct {
	ID      uuid.UUID `param:"id" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max=1000"`
}

func (r *UpdateCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Delete Comment ---
type DeleteCommentRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (r *DeleteCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
