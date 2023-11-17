package entity

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Base
	Title    string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	Caption  string `gorm:"not null; type:varchar(255)"`
	PhotoURL string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	UserID   uint   `gorm:"not null; type:int"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	// check user
	var user User
	if err := tx.Where("id = ?", p.UserID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	return nil
}
