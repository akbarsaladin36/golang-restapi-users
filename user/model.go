package user

import (
	"gorm.io/datatypes"
)

type User struct {
	UserId          int    `gorm:"primaryKey"`
	UserName        string `json:"user_username"`
	UserPassword    string `json:"user_password"`
	UserEmail       string `json:"user_email"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber int    `json:"user_phone_number"`
	UserCreatedDate datatypes.Date
	UserUpdatedDate datatypes.Date
}
