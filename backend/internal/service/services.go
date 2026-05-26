package service

import (
	"github.com/heyitsk/go-boilerplate/internal/lib/job"
	"github.com/heyitsk/go-boilerplate/internal/repository"
	"github.com/heyitsk/go-boilerplate/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}