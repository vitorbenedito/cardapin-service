package model

import (
	"errors"
	"strings"

	"cardap.in/db"
)

type Table struct {
	ID        uint     `gorm:"primary_key; type:serial; not null;"`
	Name      string   `gorm:"type:varchar(255); not null;"`
	CompanyID uint     `gorm:"type:bigint; not null;"`
	Company   *Company `gorm:"foreignkey:CompanyID"`
}

type TableJSON struct {
	ID      uint                `json:"id"`
	Name    string              `json:"name"`
	Company *CompanyMinimalJSON `json:"company"`
}

func (t Table) AsJSON() *TableJSON {
	return &TableJSON{t.ID, t.Name, t.Company.AsMinimalJSON()}
}

func (t TableJSON) AsModel() *Table {
	return &Table{t.ID, t.Name, t.Company.ID, nil}
}

func (t TableJSON) GetId() uint {
	return t.ID
}

func (Table) TableName() string {
	return "cardapin_table"
}

func (t *Table) HasConflict() (error, []string) {
	return t.hasSameTable()
}

func (t *Table) tableExists() bool {
	var tableLoaded Table
	db.DB.Where("name = ? and id != ? and company_id = ?", strings.TrimSpace(t.Name), t.ID, t.CompanyID).Find(&tableLoaded)
	return tableLoaded.ID != 0
}

func (t *Table) hasSameTable() (error, []string) {
	if t.tableExists() {
		return errors.New("table.errors.alreadyexists"), []string{t.Name}
	}
	return nil, nil
}
