package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/utils"
)

type User struct {
	gorm.Model
	Username     string        `gorm:"not null; type:varchar(255);unique" valid:"required"`
	Email        string        `gorm:"not null; type:varchar(255);unique" valid:"required, email"`
	Password     string        `gorm:"not null; type:varchar(255)"        valid:"required, minstringlength(6)"`
	Age          uint          `gorm:"not null; type:int"                 valid:"required, range(8|100)"`
	Photos       []Photo       `gorm:"foreignKey:UserID"`
	Comments     []Comment     `gorm:"foreignKey:UserID"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	// hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return nil
}
