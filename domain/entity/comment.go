package entity

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `gorm:"not null; type:varchar(255)" valid:"required"`
	UserID  uint   `gorm:"not null; type:int"`
	PhotoID uint   `gorm:"not null; type:int"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	photo := Photo{}
	if err := tx.First(&photo, c.PhotoID).Error; err != nil {
		return errors.New("photo not found")
	}

	user := User{}
	if err := tx.First(&user, c.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	return nil
}
