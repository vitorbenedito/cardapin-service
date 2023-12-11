package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Address    string `gorm:"type:varchar(255);not null;"`
	Area       string `gorm:"type:varchar(255);"`
	City       string `gorm:"type:varchar(255);"`
	State      string `gorm:"type:varchar(255);"`
	PostalCode string `gorm:"type:varchar(255);"`
	Complement string `gorm:"type:varchar(255);"`
	Number     string `gorm:"type:varchar(10);not null;"`
	Latitude   string `gorm:"type:varchar(255);"`
	Longitude  string `gorm:"type:varchar(255);"`
	CompanyID  uint   `gorm:"type:bigint;not null;"`
}

type AddressJSON struct {
	ID         uint   `json:"id"`
	Address    string `json:"address"`
	Area       string `json:"area"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Complement string `json:"complement"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Number     string `json:"number"`
}

func (o Address) asJSON() *AddressJSON {
	return &AddressJSON{o.ID, o.Address, o.Area, o.City, o.State, o.PostalCode, o.Complement, o.Latitude, o.Longitude, o.Number}
}

func (o AddressJSON) AsModel() *Address {
	return &Address{gorm.Model{ID: o.ID}, o.Address, o.Area, o.City, o.State, o.PostalCode, o.Complement, o.Number, o.Latitude, o.Longitude, 0}
}

func (Address) TableName() string {
	return "address"
}
