package service

import (
	"fmt"

	"github.com/ApoorvYdv/go-tasker/internal/lib/aws"
	"github.com/ApoorvYdv/go-tasker/internal/lib/job"
	"github.com/ApoorvYdv/go-tasker/internal/repository"
	"github.com/ApoorvYdv/go-tasker/internal/server"
)

type Services struct {
	Auth     *AuthService
	Job      *job.JobService
	Todo     *TodoService
	Comment  *CommentService
	Category *CategoryService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)
	awsClient, err := aws.NewAWS(s)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AWS client: %w", err)
	}

	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Category: NewCategoryService(s, repos.Category),
		Todo:     NewTodoService(s, repos.Todo, repos.Category, awsClient),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
	}, nil
}
