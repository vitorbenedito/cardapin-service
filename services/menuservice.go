package services

import (
	"errors"
	"strconv"

	"cardap.in/lambda/apperrors"

	"github.com/jinzhu/gorm"

	"cardap.in/lambda/db"
	"cardap.in/lambda/model"
)

type MenuServices struct {
}

func (ms *MenuServices) Save(menuToSave model.Menu) (model.MenuJSON, error) {
	if !db.DB.NewRecord(&menuToSave) {
		return ms.Update(menuToSave)
	}
	companyServices := CompanyServices{}
	companyID := strconv.FormatUint(uint64(menuToSave.CompanyID), 10)
	menuToSave.Company = companyServices.List(companyID)

	for _, category := range menuToSave.Categories {
		for _, product := range category.Products {
			groups := []*model.AdditionalItemsGroup{}
			db.DB.Where("id IN (?)", product.GetGroupsIds()).Find(&groups)
			product.AdditionalItemsGroups = groups
		}
	}
	db.DB.Save(&menuToSave)
	ms.preloadAllMenus(db.DB).Find(&menuToSave)
	return *menuToSave.AsJSON(), nil
}

func (ms *MenuServices) GetMenuEnabledByCompanyCode(companyCode string) (model.MenuJSON, error) {
	var menu model.Menu
	db.DB.Joins("LEFT JOIN company on company.id = menu.company_id").Where("company.company_code = ? AND menu.enabled = true", companyCode).Preload("Company").Preload("Company.Addresses").Preload("Company.Sections").Preload("Company.PaymentTypes").Preload("Company.OpeningHours").Preload("Categories").Preload("Categories.Products").Preload("Categories.Products.AdditionalItemsGroups").Preload("Categories.Products.AdditionalItemsGroups.AdditionalItems").Find(&menu)
	return *menu.AsJSON(), nil
}

func (ms *MenuServices) EnableMenu(menuId string) (model.MenuJSON, error) {
	var menu model.Menu
	id, _ := strconv.ParseInt(menuId, 10, 64)
	db.DB.Where("id = ?", id).Find(&menu)
	if menu.ID == 0 {
		return *&model.MenuJSON{}, errors.New("Id not found")
	}
	menu.Enabled = true
	db.DB.Save(&menu)
	db.DB.Model(&model.Menu{}).Where("id != ? AND company_id = ?", id, menu.CompanyID).UpdateColumn("enabled", false)
	if menu.ID != 0 {
		ms.preloadAllMenus(db.DB).Find(&menu)
		menu.HideDisabled()
	}
	return *menu.AsJSON(), nil
}

func (ms *MenuServices) DeleteMenu(menuId string) (bool, error) {
	id, _ := strconv.ParseInt(menuId, 10, 64)
	db.DB.Unscoped().Where("id = ?", id).Delete(&model.Menu{})
	return true, nil
}

func (ms *MenuServices) GetMenuByLoggedCompany(token string) ([]*model.MenuJSON, *apperrors.AppError) {

	if len(token) == 0 {
		return []*model.MenuJSON{}, &apperrors.AppError{errors.New("Logged company not found"), "", 404}
	}

	companyService := CompanyServices{}
	company, err := companyService.GetCompanyByToken(token)
	if err != nil {
		return nil, err
	}

	var menus []model.Menu
	query := db.DB.Joins("LEFT JOIN company on company.id = menu.company_id").Where("company.company_code = ?", company.CompanyCode)
	ms.preloadAllMenus(query).Find(&menus)

	menuJSON := make([]*model.MenuJSON, 0)
	for _, menu := range menus {
		menuJSON = append(menuJSON, menu.AsJSON())
	}
	return menuJSON, nil
}

func (ms *MenuServices) Update(menuToSave model.Menu) (model.MenuJSON, error) {
	companyServices := CompanyServices{}
	companyID := strconv.FormatUint(uint64(menuToSave.CompanyID), 10)
	menuToSave.Company = companyServices.List(companyID)

	menuLoaded := loadByID(menuToSave.ID)

	categories, products := menuToSave.GetAssociationIDs()
	categoriesToRemove, productsToRemove := menuLoaded.GetNotUsedIds(categories, products)
	db.DB.Unscoped().Where("id IN (?)", productsToRemove).Delete(model.Product{})
	db.DB.Unscoped().Where("id IN (?)", categoriesToRemove).Delete(model.Category{})
	for _, category := range menuToSave.Categories {
		for _, product := range category.Products {
			ids := product.GetGroupsIds()
			if len(ids) == 0 {
				db.DB.Exec("DELETE FROM product_additional_group WHERE product_id = ?", product.ID)
				continue
			}
			groups := []*model.AdditionalItemsGroup{}
			db.DB.Where("id IN (?)", ids).Find(&groups)
			product.AdditionalItemsGroups = groups
			db.DB.Exec("DELETE FROM product_additional_group WHERE product_id = ? AND additional_items_group_id NOT IN (?)", product.ID, ids)
		}
	}
	db.DB.Save(&menuToSave)
	db.DB.Preload("Company").Preload("Company.Addresses").Preload("Company.Sections").Preload("Company.PaymentTypes").Preload("Company.OpeningHours").Preload("Categories").Preload("Categories.Products").Preload("Categories.Products.AdditionalItemsGroups").Preload("Categories.Products.AdditionalItemsGroups.AdditionalItems").Where("id = ?", menuToSave.ID).Find(&menuToSave)
	return *menuToSave.AsJSON(), nil
}

func loadByID(id uint) model.Menu {
	loaded := model.Menu{}
	db.DB.Where("id = ? ", id).Preload("Categories").Preload("Categories.Products").Preload("Categories.Products.AdditionalItemsGroups").Preload("Categories.Products.AdditionalItemsGroups.AdditionalItems").Find(&loaded)
	return loaded
}

func (ms *MenuServices) preloadAllMenus(db2 *gorm.DB) *gorm.DB {
	return db2.Preload("Company").Preload("Company.Addresses").Preload("Company.Sections").Preload("Company.PaymentTypes").Preload("Company.OpeningHours").Preload("Categories").Preload("Categories.Products").Preload("Categories.Products.AdditionalItemsGroups").Preload("Categories.Products.AdditionalItemsGroups.AdditionalItems")
}
