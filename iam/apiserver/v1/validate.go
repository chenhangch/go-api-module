package v1

import "time"

type User struct {
	Status int `json:"status" gorm:"column:status" validate:"omitempty"`

	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`

	Password string `json:"password" gorm:"column:password" validate:"required"`

	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"_" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

func (u *User) TableName() string {
	return "user"
}

