package user

import (
	"time"

	"gorm.io/datatypes"
)

type Service interface {
	GetUsers() ([]User, error)
	GetUser(ID int) (User, error)
	CreateUser(userInput UserInput) (User, error)
	UpdateUser(ID int, userInput UserInput) (User, error)
	DeleteUser(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUsers() ([]User, error) {
	users, err := s.repository.GetAll()

	return users, err
}

func (s *service) GetUser(ID int) (User, error) {
	user, err := s.repository.GetID(ID)

	return user, err
}

func (s *service) CreateUser(userInput UserInput) (User, error) {

	user := User{
		UserName:        userInput.UserName,
		UserPassword:    userInput.UserPassword,
		UserEmail:       userInput.UserEmail,
		UserAddress:     userInput.UserAddress,
		UserPhoneNumber: userInput.UserPhoneNumber,
		UserCreatedDate: datatypes.Date(time.Now()),
	}

	newUser, err := s.repository.Create(user)

	return newUser, err
}

func (s *service) UpdateUser(ID int, userInput UserInput) (User, error) {
	checkUser, _ := s.repository.GetID(ID)

	checkUser.UserName = userInput.UserName
	checkUser.UserPassword = userInput.UserPassword
	checkUser.UserEmail = userInput.UserEmail
	checkUser.UserAddress = userInput.UserAddress
	checkUser.UserPhoneNumber = userInput.UserPhoneNumber
	checkUser.UserUpdatedDate = datatypes.Date(time.Now())

	updateUser, err := s.repository.Update(checkUser)

	return updateUser, err
}

func (s *service) DeleteUser(ID int) (User, error) {
	checkUser, _ := s.repository.GetID(ID)

	deleteUser, err := s.repository.Delete(checkUser)

	return deleteUser, err
}
