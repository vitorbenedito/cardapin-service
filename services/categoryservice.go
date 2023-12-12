package services

import (
	"cardap.in/db"
	"cardap.in/model"
)

type CategoryServices struct {
}

func (*CategoryServices) SaveOrUpdate(categoryToSave model.Category) (model.CategoryJSON, error) {
	db.DB.Save(&categoryToSave)
	db.DB.Preload("Company").Find(&categoryToSave)
	return *categoryToSave.AsJSON(), nil
}

func (*CategoryServices) GetByCompanyId(id string) (*[]model.CategoryJSON, error) {

	var categories *[]model.Category
	db.DB.Preload("Company").Where("company_id = ?", id).Find(&categories)
	categoriesJson := make([]model.CategoryJSON, len(*categories))
	for _, category := range *categories {
		categoriesJson = append(categoriesJson, *category.AsJSON())
	}
	return &categoriesJson, nil
}
