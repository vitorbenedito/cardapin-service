package services

import (
	"errors"

	"cardap.in/lambda/db"
	"cardap.in/lambda/model"
)

type UserService struct {
}

func (*UserService) SaveUser(userToSave model.User) (model.UserJSON, error) {
	if userExists(userToSave) {
		return model.UserJSON{}, errors.New("User with login " + userToSave.Login + " already exists")
	}
	companyServices := &CompanyServices{}
	if companyServices.CompanyCodeExists(*userToSave.Company) {
		return model.UserJSON{}, errors.New("Company code with name " + userToSave.Company.CompanyCode + " already exists")
	}

	db.DB.Save(&userToSave)
	return userToSave.AsJSON(), nil
}

func (*UserService) GetUserById(id string) (model.UserJSON, error) {
	user := model.User{}
	db.DB.Preload("Companies").Where("ID = ?", id).Find(&user)
	return user.AsJSON(), nil
}

func (*UserService) Login(userToSave model.User) (model.UserJSON, error) {
	user := logUser(userToSave)
	if user.ID == 0 {
		return model.UserJSON{}, errors.New("Username or password wrong")
	}
	return user.AsJSON(), nil
}

func userExists(userToSave model.User) bool {
	var userFromDB model.User
	db.DB.Raw("SELECT * FROM cardapin_user WHERE login = ? and id != ?", userToSave.Login, userToSave.ID).Scan(&userFromDB)
	return userFromDB.ID != 0
}

func logUser(userToValidate model.User) model.User {
	var userFromDB model.User
	db.DB.Where("login = ? AND password = ?", userToValidate.Login, userToValidate.Password).Preload("Company").Preload("Company.Sections").Preload("Company.Addresses").Preload("Company.PaymentTypes").Preload("Company.OpeningHours").Find(&userFromDB)
	return userFromDB
}
