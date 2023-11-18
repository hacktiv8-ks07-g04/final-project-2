package entity

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Base
	UserID  uint   `gorm:"not null; type:int"`
	PhotoID uint   `gorm:"not null; type:int"`
	Message string `gorm:"not null; type:varchar(255)"                    valid:"required"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
	Photo   Photo  `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	// check photo exist
	var photo Photo
	if err := tx.First(&photo, c.PhotoID).Error; err != nil {
		return errors.New("photo not found")
	}

	return nil
}
