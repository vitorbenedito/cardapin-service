package model

import (
	"errors"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Description  string `gorm:"type:varchar(255);"`
	Enabled      bool   `gorm:"type:boolean;"`
	MinimumOrder int64  `gorm:"type:bigint;"`
	CompanyID    uint   `gorm:"type:bigint;"`
	Company      Company
	Categories   []*Category
}

func (Menu) TableName() string {
	return "menu"
}

func (m *Menu) HideDisabled() {
	categoriesToShow := make([]*Category, 0)
	for _, category := range m.Categories {
		category.HideDisabled()
		if category.Enabled {
			categoriesToShow = append(categoriesToShow, category)
		}
	}
	m.Categories = categoriesToShow
}
func (m *Menu) HasConflict() (error, []string) {
	return m.hasSameCategoriesOrProducts(), nil
}

func (m *Menu) GetAssociationIDs() ([]uint, []uint) {
	categories := make([]uint, 0)
	products := make([]uint, 0)
	for _, category := range m.Categories {
		categories = append(categories, category.ID)
		for _, product := range category.Products {
			products = append(products, product.ID)
		}
	}
	return categories, products
}

func (m *Menu) GetNotUsedIds(otherCategories []uint, otherProducts []uint) ([]uint, []uint) {
	categories, products := m.GetAssociationIDs()
	return getNotFoundValues(categories, otherCategories), getNotFoundValues(products, otherProducts)
}

func (m *Menu) hasSameCategoriesOrProducts() error {
	names := map[string]bool{}
	for _, c := range m.Categories {
		names[c.Name] = true
		if err := c.hasSameProducts(); err != nil {
			return err
		}
	}
	if len(m.Categories) != len(names) {
		return errors.New("category.errors.alreadyexists")
	}
	return nil
}

type MenuJSON struct {
	ID           uint            `json:"id"`
	Description  string          `json:"description"`
	Enabled      bool            `json:"enabled"`
	MinimumOrder int64           `json:"minimumOrder"`
	Company      *CompanyJson    `json:"company"`
	Categories   []*CategoryJSON `json:"categories"`
}

func (m MenuJSON) GetId() uint {
	return m.ID
}

func (m *Menu) AsJSON() *MenuJSON {
	menuJSON := &MenuJSON{
		ID:           m.ID,
		Description:  m.Description,
		Enabled:      m.Enabled,
		MinimumOrder: m.MinimumOrder,
		Company:      m.Company.AsJson(),
		Categories:   make([]*CategoryJSON, 0),
	}
	categoriesToAdd := make([]*CategoryJSON, 0)
	for _, category := range m.Categories {
		categoriesToAdd = append(categoriesToAdd, category.AsJSON())
	}
	menuJSON.Categories = categoriesToAdd
	return menuJSON
}

func (m *MenuJSON) AsModel() *Menu {
	menu := &Menu{
		gorm.Model{ID: m.ID},
		m.Description,
		m.Enabled,
		m.MinimumOrder,
		m.Company.ID,
		Company{},
		make([]*Category, 0),
	}
	categoriesToAdd := make([]*Category, 0)
	for _, category := range m.Categories {
		categoriesToAdd = append(categoriesToAdd, category.AsModel())
	}
	menu.Categories = categoriesToAdd
	return menu
}
