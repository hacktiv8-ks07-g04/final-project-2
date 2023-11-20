package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string    `gorm:"not null"                                      valid:"required"`
	Caption  string    `gorm:"not null"`
	PhotoURL string    `gorm:"not null"                                      valid:"required"`
	UserID   uint      `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
