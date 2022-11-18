package models

import (
	"github.com/jinzhu/gorm"
)

type UserExcel struct {
	gorm.Model
	Empcode      string `gorm:"size:255;not null;" json:"empcode"`
	Branch       string `gorm:"size:255;not null;" json:"branch"`
	Role         string `gorm:"size:255;null;" json:"role"`
	Moblienumber string `gorm:"size:255;not null;" json:"moblienumber"`
	Emailid      string `gorm:"size:255;not null;" json:"emailid"`
	Usertpe      string `gorm:"size:255;not null;" json:"usertype"`
}

func (u *UserExcel) SaveData() (*UserExcel, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &UserExcel{}, err
	}
	return u, nil
}
