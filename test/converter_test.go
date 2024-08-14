package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/karimatiket/karimatiket-user/internal/domain/user"
	"github.com/karimatiket/karimatiket-user/internal/infrastructure/db/postgres"
	"github.com/karimatiket/karimatiket-user/internal/infrastructure/validation"
	"github.com/karimatiket/karimatiket-user/internal/pb"
)

func TestCreateUser(t *testing.T) {
	godotenv.Load()

	pg, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	pg.AutoMigrate(&user.User{})

	v := validation.New()
	s := user.NewService(pg, v)

	pbUserCreateRequest := &pb.UserCreateRequest{
		IdentityType:   "KTP",
		IdentityNumber: "1234567890123456",
		FullName:       "Agil Ghani Istikmal",
		Email:          "agil_g@safatanc.com",
		Phone:          "+6281234567890",
		Address:        "Mempawah, Kalimantan Barat",
		BirthDate:      "2004-10-21",
	}

	pbUserResponse, err := s.CreateUser(pbUserCreateRequest)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(pbUserResponse)
}
