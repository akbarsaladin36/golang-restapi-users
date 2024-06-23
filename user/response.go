package user

type UserResponse struct {
	UserID          int    `json:"user_id"`
	UserName        string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber int    `json:"user_phone_number"`
}
