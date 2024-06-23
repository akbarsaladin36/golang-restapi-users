package user

type UserInput struct {
	UserName        string `json:"user_username"`
	UserPassword    string `json:"user_password"`
	UserEmail       string `json:"user_email"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber int    `json:"user_phone_number"`
}
