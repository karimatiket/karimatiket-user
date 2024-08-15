package test

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/karimatiket/karimatiket-user/internal/domain/user"
	"github.com/karimatiket/karimatiket-user/internal/infrastructure/db/postgres"
	"github.com/karimatiket/karimatiket-user/internal/infrastructure/validation"
	"github.com/karimatiket/karimatiket-user/pb"
	"github.com/karimatiket/karimatiket-user/pkg/converter"
)

func TestConvertUser(t *testing.T) {
	godotenv.Load()

	birthDate, _ := time.Parse(time.DateOnly, "2004-10-21")
	input := &pb.UserCreateRequest{
		IdentityType:   "KTP",
		IdentityNumber: "1234567890123456",
		FullName:       "Agil Ghani Istikmal",
		Email:          "agil_g@safatanc.com",
		Phone:          "+6281234567890",
		Address:        "Mempawah, Kalimantan Barat",
		BirthDate:      birthDate.Format(time.RFC3339),
	}

	user, err := converter.TypeConverter[user.User](input)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(user)
}

func TestCreateUser(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	birthDate, _ := time.Parse(time.DateOnly, "2004-10-21")
	pbUserCreateRequest := &pb.UserCreateRequest{
		IdentityType:   "KTP",
		IdentityNumber: "6102011223457773",
		FullName:       "Gronw Amunet",
		Email:          "gronw_a@safatanc.com",
		Phone:          "+6281234567891",
		Address:        "Greek Somewhere",
		BirthDate:      birthDate.Format(time.RFC3339),
	}

	pbUserResponse, err := s.CreateUser(pbUserCreateRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(pbUserResponse)
}

func TestUpdateUser(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	address := "Kota Pontianak"
	pbUserUpdateRequest := &pb.UserUpdateRequest{
		Id: "a9fd612a-b92b-49c7-b6a0-70093d513920",
		Data: &pb.UserUpdateRequestData{
			Address: &address,
		},
	}

	pbUserResponse, err := s.UpdateUser(pbUserUpdateRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(pbUserResponse)
}

func TestGetUsers(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	pbUserResponses, err := s.GetUsers()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, user := range pbUserResponses.Users {
		log.Println(user)
	}
}

func TestGetUser(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	pbUserResponses, err := s.GetUser(&pb.UserGetRequest{
		Id: "a9fd612a-b92b-49c7-b6a0-70093d513920",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(pbUserResponses)
}

func TestDeleteUser(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	pbUserResponses, err := s.DeleteUser(&pb.UserDeleteRequest{
		Id: "cfa54b4d-30dc-498f-bb29-8a6202aba3e7",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(pbUserResponses)
}
