package services

import (
	"errors"
	"strconv"

	"cardap.in/db"
	"cardap.in/model"
)

type TableServices struct {
}

func (t *TableServices) Save(tableToSave model.Table, isUpdate bool) (*model.Table, error) {
	if !isUpdate && !db.DB.NewRecord(&tableToSave) {
		return nil, errors.New("Could not update a request with a POST")
	}

	companyServices := CompanyServices{}
	companyID := strconv.FormatUint(uint64(tableToSave.CompanyID), 10)
	company := companyServices.List(companyID)
	tableToSave.Company = &company
	db.DB.Save(&tableToSave)
	t.preload(&tableToSave)
	return &tableToSave, nil
}

func (*TableServices) Delete(tableId string) (bool, error) {
	id, _ := strconv.ParseInt(tableId, 10, 64)
	db.DB.Unscoped().Where("id = ?", id).Delete(&model.Table{})
	return true, nil
}

func (t *TableServices) List(companyID string) ([]*model.TableJSON, error) {
	companyIDuint, _ := strconv.ParseInt(companyID, 10, 64)
	var tables []*model.Table
	db.DB.Joins("LEFT JOIN company on company.id = cardapin_table.company_id").Where("company.id = ?", companyIDuint).Preload("Company").Find(&tables)
	tablesJSON := make([]*model.TableJSON, 0)
	for _, table := range tables {
		tablesJSON = append(tablesJSON, table.AsJSON())
	}
	return tablesJSON, nil
}

func (*TableServices) preload(table *model.Table) {
	db.DB.Preload("Company").Find(&table)
}
