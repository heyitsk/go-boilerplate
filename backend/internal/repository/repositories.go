package repository


import "github.com/heyitsk/go-boilerplate/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}