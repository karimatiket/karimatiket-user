package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid()" json:"id,omitempty"`
	IdentityType   string     `validate:"required,uppercase,min=3" json:"identity_type,omitempty"`
	IdentityNumber string     `validate:"required,number,len=16" gorm:"unique" json:"identity_number,omitempty"`
	FullName       string     `validate:"required,min=3" json:"full_name,omitempty"`
	Email          string     `validate:"required,email" gorm:"unique" json:"email,omitempty"`
	Phone          string     `validate:"required,e164" gorm:"unique" json:"phone,omitempty"`
	Address        string     `validate:"required,min=8" json:"address,omitempty"`
	BirthDate      *time.Time `validate:"required" json:"birth_date,omitempty"`
	CreatedAt      *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

type UserUpdateValidation struct {
	IdentityType   string     `validate:"omitempty,uppercase,min=3" json:"identity_type,omitempty"`
	IdentityNumber string     `validate:"omitempty,number,len=16" gorm:"unique" json:"identity_number,omitempty"`
	FullName       string     `validate:"omitempty,min=3" json:"full_name,omitempty"`
	Email          string     `validate:"omitempty,email" gorm:"unique" json:"email,omitempty"`
	Phone          string     `validate:"omitempty,e164" gorm:"unique" json:"phone,omitempty"`
	Address        string     `validate:"omitempty,min=8" json:"address,omitempty"`
	BirthDate      *time.Time `validate:"omitempty" json:"birth_date,omitempty"`
}

type UserIdValidation struct {
	Id string `validate:"required,uuid"`
}
