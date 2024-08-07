package user

import (
	"github.com/karimatiket/karimatiket-user/pb"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) CreateUser(userCreateRequest *pb.UserCreateRequest) (*pb.User, error) {
	return nil, nil
}
