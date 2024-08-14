package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid()"`
	IdentityType   string     `validate:"required,uppercase,min=3"`
	IdentityNumber string     `validate:"required,number,len=16" gorm:"unique"`
	FullName       string     `validate:"required,min=3"`
	Email          string     `validate:"required,email" gorm:"unique"`
	Phone          string     `validate:"required,e164" gorm:"unique"`
	Address        string     `validate:"required,min=8"`
	BirthDate      *time.Time `validate:"required"`
	CreatedAt      *time.Time `gorm:"autoCreateTime"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime"`
}
