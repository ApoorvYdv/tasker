package comment

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// --- Add Comment ---
type AddCommentPayload struct {
	TodoID  uuid.UUID `json:"todoId" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max=1000"`
}

func (r *AddCommentPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Get Comments ---
type GetCommentsByTodoIDPayload struct {
	TodoID uuid.UUID `param:"todoId" validate:"required,uuid"`
}

func (r *GetCommentsByTodoIDPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Update Comment ---
type UpdateCommentPayload struct {
	ID      uuid.UUID `param:"id" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max=1000"`
}

func (r *UpdateCommentPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// --- Delete Comment ---
type DeleteCommentPayload struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (r *DeleteCommentPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
