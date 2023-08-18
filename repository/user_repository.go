package repository

import (
	"fmt"
	"staycation/domain"

	"gorm.io/gorm"
)

type UserRepositoryContract interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(ID string) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}
type UserRepository struct {
  db *gorm.DB
}
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db:db}
}

func (r UserRepository) Save(users domain.User) (domain.User, error){
	err := r.db.Create(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
func (r UserRepository) FindByID(ID string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id =?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r UserRepository) Login(ID string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id =?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	fmt.Println(email)
	err := r.db.Where("email =?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	
	fmt.Println(user.ID)
	return user, nil
}
func (r UserRepository) Update(user domain.User) (domain.User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}