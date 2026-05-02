package userservice

import userrepo "steam-tracker/internal/repository/user"

type Service struct {
	repo userrepo.UserRepository
}

func New(repo userrepo.UserRepository) *Service {
	return &Service{repo: repo}
}
