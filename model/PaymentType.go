package model

import "github.com/jinzhu/gorm"

type PaymentType struct {
	gorm.Model
	Name    string     `gorm:"type:varchar(255);not null;"`
	Company []*Company `gorm:"many2many:company_payment_type"`
}

type PaymentTypeJSON struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s PaymentType) AsJSON() *PaymentTypeJSON {
	return &PaymentTypeJSON{s.ID, s.Name}
}

func (s PaymentTypeJSON) AsModel() *PaymentType {
	return &PaymentType{gorm.Model{ID: s.ID},
		s.Name, make([]*Company, 0)}
}

func (PaymentType) TableName() string {
	return "payment_type"
}
