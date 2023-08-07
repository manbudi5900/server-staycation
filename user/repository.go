package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID string) (User, error)
	Update(user User) (User, error)
}
type repository struct {
  db *gorm.DB
}
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(users User) (User, error){
	err := r.db.Create(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
func (r *repository) FindByID(ID string) (User, error) {
	var user User
	err := r.db.Where("id =?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) Login(ID string) (User, error) {
	var user User
	err := r.db.Where("id =?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	fmt.Println(email)
	err := r.db.Where("email =?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	
	fmt.Println(user.ID)
	return user, nil
}
func (r *repository) Update(user User) (User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}