package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)
	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return err
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)
	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return err
}
