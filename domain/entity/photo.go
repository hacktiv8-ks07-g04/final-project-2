package entity

import (
	"github.com/asaskevich/govalidator"
)

type Photo struct {
	Base
	Title    string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	Caption  string `gorm:"not null; type:varchar(255)"`
	PhotoURL string `gorm:"not null; type:varchar(255)"                   valid:"required"`
	UserID   uint   `gorm:"not null; type:int"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}

func (p *Photo) BeforeCreate() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
