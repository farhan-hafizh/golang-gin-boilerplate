package user

import "gorm.io/gorm"

// used for interacting with service
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

//private
type repository struct {
	//variable with type gorm.DB
	db *gorm.DB
}

// constructor
func CreateRepository(db *gorm.DB) *repository {
	// return the repository where db is db from params
	return &repository{db}
}

// create a function called save for "repository" that the
// parameter is user object and the return is User object or error
func (r *repository) Save(user User) (User, error) {
	//create user object on db with user data from params
	//and return assign error to err if error
	err := r.db.Create(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	// init variable with type user
	var user User
	// find in table user where email = email and save it to variable user
	err := r.db.Where("email = ?", email).Find(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) FindById(id int) (User, error) {
	// init variable with type user
	var user User
	// find in table user where id = id and save it to variable user
	err := r.db.Where("id = ?", id).Find(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	// update user data in db that have the parameter user's id
	err := r.db.Save(&user).Error

	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}
