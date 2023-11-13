package entity

import (
	"github.com/asaskevich/govalidator"
)

type User struct {
	Base
	Username string `gorm:"not null; type:varchar(255);unique" valid:"required"`
	Email    string `gorm:"not null; type:varchar(255);unique" valid:"required, email"`
	Password string `gorm:"not null; type:varchar(255)"        valid:"required, minstringlength(6)"`
	Age      uint   `gorm:"not null; type:int"                 valid:"required, range(8|100)"`
}

func (u *User) BeforeCreate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}
