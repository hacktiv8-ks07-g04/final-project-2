package entity

import (
	"github.com/asaskevich/govalidator"
)

type Comment struct {
	Base
	UserID  uint   `gorm:"not null; type:int"`
	PhotoID uint   `gorm:"not null; type:int"`
	Message string `gorm:"not null; type:varchar(255)"                    valid:"required"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
	Photo   Photo  `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE"`
}

func (c *Comment) BeforeCreate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
