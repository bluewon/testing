package user

import (
	"crypto/sha1"
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	h := sha1.New()
	h.Write([]byte(req.Pwd + user.Salt))
	hash := h.Sum(nil)
	if fmt.Sprintf("%x", hash) != user.Hash {
		return nil, fmt.Errorf("password incorrect")
	}
	return &LoginResponse{
		UserName: user.Name,
	}, nil
}
