package repository

import "tagus/model"

type UserRepository struct {
	Repository
}

func (r UserRepository) Find(column []string, q model.User) (model.User, error) {
	var u model.User

	tx := r.DB.Where(q)

	if len(column) != 0 {
		tx.Select(column)
	}

	err := tx.First(&u).Error

	return u, err
}

func (r UserRepository) Create(userName, password, displayName string) (model.User, error) {
	u := model.User{UserName: userName, Password: password, DisplayName: displayName}

	err := r.DB.Create(&u).Error

	return u, err
}
