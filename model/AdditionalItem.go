package model

import "github.com/jinzhu/gorm"

type AdditionalItem struct {
	gorm.Model
	Name                   string `gorm:"type:varchar(255);not null;"`
	Price                  int64  `gorm:"type:bigint;not null;"`
	Description            string `gorm:"type:varchar(255);not null;"`
	Enabled                bool   `gorm:"type:boolean;not null;"`
	AdditionalItemsGroupID uint   `gorm:"type:bigint;not null;"`
}

func (AdditionalItem) TableName() string {
	return "additional_item"
}

type AdditionalItemJSON struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

func (o AdditionalItem) asJSON() *AdditionalItemJSON {
	return &AdditionalItemJSON{o.ID, o.Name, o.Price, o.Description, o.Enabled}
}

func (o AdditionalItemJSON) AsModel() *AdditionalItem {
	return &AdditionalItem{gorm.Model{ID: o.ID},
		o.Name, o.Price, o.Description, o.Enabled, 0}
}
