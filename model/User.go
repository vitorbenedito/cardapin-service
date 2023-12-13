package model

import (
	"crypto/md5"
	"fmt"

	"gorm.io/gorm"
)

const (
	CardapinAdmin = "cardapin-admin"
)

type User struct {
	gorm.Model
	Name      string   `gorm:"type:varchar(255); not null;"`
	Login     string   `gorm:"type:varchar(255); not null;"`
	Email     string   `gorm:"type:varchar(255); not null;"`
	Password  string   `gorm:"type:varchar(255); not null;"`
	CompanyID uint     `gorm:"type:bigint;"`
	Company   *Company `gorm:"foreignKey:CompanyID"`
}

type UserRequestJSON struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Login    string       `json:"login"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Company  *CompanyJson `json:"company"`
}

type UserJSON struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Login   string       `json:"login"`
	Email   string       `json:"email"`
	Company *CompanyJson `json:"company"`
}

type UserLoginJSON struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequestJSON) AsModel() User {
	encodedPass := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	var company *Company
	companyID := uint(0)
	if u.Company != nil {
		company = u.Company.AsModel()
		companyID = company.ID
	} else {
		company = nil
	}
	user := &User{
		gorm.Model{ID: u.ID},
		u.Name,
		u.Login,
		u.Email,
		encodedPass,
		companyID,
		company,
	}
	return *user
}

func (u *UserLoginJSON) AsModel() User {
	encodedPass := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	user := &User{
		Login:    u.Login,
		Password: encodedPass,
		Email:    u.Email,
	}
	return *user
}

func (u *UserJSON) IsAdmin() bool {
	return u.Login == CardapinAdmin
}

func (u *User) AsJSON() UserJSON {
	var company *CompanyJson
	if u.Company != nil {
		company = u.Company.AsJson()
	} else {
		company = nil
	}
	userJSON := &UserJSON{
		u.ID,
		u.Name,
		u.Login,
		u.Email,
		company,
	}
	return *userJSON
}

func (User) TableName() string {
	return "cardapin_user"
}

func (u UserJSON) GetId() uint {
	return u.ID
}
