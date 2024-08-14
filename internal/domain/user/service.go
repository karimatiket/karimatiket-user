package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/karimatiket/karimatiket-user/internal/pb"
	"gorm.io/gorm"
)

type service struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewService(db *gorm.DB, validate *validator.Validate) *service {
	return &service{
		DB:       db,
		Validate: validate,
	}
}

func (s *service) CreateUser(input *pb.UserCreateRequest) (*pb.UserResponse, error) {
	user, err := PbUserCreateRequestToUser(input)
	if err != nil {
		return nil, err
	}

	err = s.Validate.Struct(user)
	if err != nil {
		return nil, err
	}

	err = s.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	pbUserResponse := UserToPbUserResponse(user)

	return pbUserResponse, nil
}

func (s *service) UpdateUser(input *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (s *service) DeleteUser(input *pb.UserDeleteRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (s *service) GetUser(input *pb.UserGetRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (s *service) GetUsers() *pb.UserResponse {
	var users []*User
	s.DB.Find(&users)
	return nil
}
