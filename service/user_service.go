package service

import (
	"errors"
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceContract interface {
  RegisterUser(input dto.RegisterUserInput) (domain.User, error)
  LoginUser(input dto.LoginUserInput) (domain.User, error)
  GetUserByID(ID string) (domain.User, error)
  SaveAvatar(ID string, fileLocation string) (domain.User, error)


}
type UserService struct {
	UserRepository repository.UserRepository
}
func NewUserService(m repository.UserRepository) UserService{
	return UserService{m}
}
func (s UserService) RegisterUser(input dto.RegisterUserInput) (domain.User, error){
	user := domain.User{}
	
	user.Email = input.Email
	user.Role = "user"
	user.Name = input.Name
	user.Username = input.Username
	user.Phone = input.Phone
	user.Province = input.Province
	user.City = input.City
	user.ID = uuid.NewString()

	userCek, err := s.UserRepository.FindByEmail(input.Email)
	if userCek.ID != ""{
		return userCek, errors.New("Email sudah digunakan")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil{
		return user, err
	}
	user.Password = string(passwordHash)
	fmt.Println(user)
	newUser, err := s.UserRepository.Save(user)
	if err != nil{
		return newUser, err
	}
	return newUser, nil
}

func (s UserService) GetUserByID(ID string) (domain.User, error) {
	user, err := s.UserRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == "0" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
func (s UserService) LoginUser(input dto.LoginUserInput) (domain.User, error){
	
	email := input.Email
	password := input.Password

	user, err := s.UserRepository.FindByEmail(email)
	fmt.Println(user)
	if err != nil {
		return user, errors.New("No user found on that email")
	}
	if user.ID == "0" {
		return user, errors.New("No user found on that email")
	}
	

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("Password Salah")
	}

	return user, nil
}
func (s UserService) SaveAvatar(ID string, fileLocation string) (domain.User, error) {
	user, err := s.UserRepository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.Avatar = fileLocation
	updateUser, err := s.UserRepository.Update(user)

	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}


