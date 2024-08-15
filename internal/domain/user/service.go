package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/karimatiket/karimatiket-user/pb"
	"github.com/karimatiket/karimatiket-user/pkg/converter"
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
	user, err := converter.TypeConverter[User](input)
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

	pbUserResponse, err := converter.TypeConverter[pb.UserResponse](user)
	if err != nil {
		return nil, err
	}

	return pbUserResponse, nil
}

func (s *service) UpdateUser(input *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	user, err := converter.TypeConverter[User](input.Data)
	if err != nil {
		return nil, err
	}

	userIdValidation, err := converter.TypeConverter[UserIdValidation](input)
	if err != nil {
		return nil, err
	}

	err = s.Validate.Struct(userIdValidation)
	if err != nil {
		return nil, err
	}

	userUpdateValidation, err := converter.TypeConverter[UserUpdateValidation](input.Data)
	if err != nil {
		return nil, err
	}

	err = s.Validate.Struct(userUpdateValidation)
	if err != nil {
		return nil, err
	}

	err = s.DB.Where("id = ?", input.Id).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	pbUserResponse, err := s.GetUser(&pb.UserGetRequest{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}

	return pbUserResponse, nil
}

func (s *service) DeleteUser(input *pb.UserDeleteRequest) (*pb.UserResponse, error) {
	userIdValidation, err := converter.TypeConverter[UserIdValidation](input)
	if err != nil {
		return nil, err
	}

	err = s.Validate.Struct(userIdValidation)
	if err != nil {
		return nil, err
	}

	pbUserResponse, err := s.GetUser(&pb.UserGetRequest{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}

	user, err := converter.TypeConverter[User](pbUserResponse)
	if err != nil {
		return nil, err
	}

	err = s.DB.Delete(&user, "id = ?", input.Id).Error
	if err != nil {
		return nil, err
	}
	return pbUserResponse, nil
}

func (s *service) GetUser(input *pb.UserGetRequest) (*pb.UserResponse, error) {
	userIdValidation, err := converter.TypeConverter[UserIdValidation](input)
	if err != nil {
		return nil, err
	}

	err = s.Validate.Struct(userIdValidation)
	if err != nil {
		return nil, err
	}

	var user *User
	err = s.DB.Take(&user, "id = ?", input.Id).Error
	if err != nil {
		return nil, err
	}
	pbUserResponse, err := converter.TypeConverter[pb.UserResponse](user)
	return pbUserResponse, nil
}

func (s *service) GetUsers() (*pb.UserResponses, error) {
	var users []*User
	err := s.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var pbUserResponses pb.UserResponses
	for _, user := range users {
		pbUserResponse, err := converter.TypeConverter[pb.UserResponse](user)
		if err != nil {
			return nil, err
		}
		pbUserResponses.Users = append(pbUserResponses.Users, pbUserResponse)
	}
	return &pbUserResponses, nil
}
