package entity

import (
	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	Base
	Name           string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	SocialMediaURL string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	UserID         uint   `gorm:"not null; type:int"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}

func (sm *SocialMedia) BeforeCreate() error {
	_, err := govalidator.ValidateStruct(sm)
	if err != nil {
		return err
	}

	return nil
}
