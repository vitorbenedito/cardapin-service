package model

import (
	"errors"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null;"`
	Description string `gorm:"type:varchar(255);"`
	Order       int    `gorm:"type:int;not null;"`
	Enabled     bool   `gorm:"type:boolean;not null;"`
	Products    []*Product
	MenuID      uint `type:bigint; not null;"`
}

type CategoryJSON struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
	Enabled     bool           `json:"enabled"`
	Products    []*ProductJSON `json:"products"`
}

func (c *CategoryJSON) AsModel() *Category {
	products := make([]*Product, 0)
	for _, productJSON := range c.Products {
		products = append(products, productJSON.AsModel(c.ID))
	}
	category := &Category{
		gorm.Model{ID: c.ID},
		c.Name,
		c.Description,
		c.Order,
		c.Enabled,
		products,
		0,
	}
	return category
}

func (c *Category) AsJSON() *CategoryJSON {
	productsJSON := make([]*ProductJSON, 0)
	for _, product := range c.Products {
		productsJSON = append(productsJSON, product.AsJSON())
	}
	categoryJSON := &CategoryJSON{
		c.ID,
		c.Name,
		c.Description,
		c.Order,
		c.Enabled,
		productsJSON,
	}
	return categoryJSON
}

func (c CategoryJSON) GetId() uint {
	return c.ID
}

func (c *Category) hasSameProducts() error {
	names := map[string]bool{}
	for _, p := range c.Products {
		names[p.Name] = true
		if err := p.hasSameAddionalItem(); err != nil {
			return err
		}
	}
	if len(c.Products) != len(names) {
		return errors.New("product.errors.alreadyexists")
	}
	return nil
}

func (c *Category) HideDisabled() {
	productsToShow := make([]*Product, 0)
	for _, product := range c.Products {
		product.HideDisabled()
		if product.Enabled {
			productsToShow = append(productsToShow, product)
		}
	}
	c.Products = productsToShow
}

func (Category) TableName() string {
	return "category"
}
