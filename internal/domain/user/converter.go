package user

import (
	"encoding/json"
	"log"
	"time"

	"github.com/karimatiket/karimatiket-user/internal/pb"
)

func PbUserCreateRequestToUser(u *pb.UserCreateRequest) (*User, error) {
	birthDate, err := time.Parse(time.DateOnly, u.BirthDate)
	if err != nil {
		return nil, err
	}
	user := &User{
		IdentityType:   u.IdentityType,
		IdentityNumber: u.IdentityNumber,
		FullName:       u.FullName,
		Email:          u.Email,
		Phone:          u.Phone,
		Address:        u.Address,
		BirthDate:      &birthDate,
	}
	return user, nil
}

func UserToPbUserResponse(u *User) *pb.UserResponse {
	jsonUser, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	var pbUserResponse *pb.UserResponse
	err = json.Unmarshal(jsonUser, &pbUserResponse)
	if err != nil {
		log.Fatal(err)
	}
	return pbUserResponse
}
