package model

import (
	"errors"
	"strconv"
	"strings"

	"cardap.in/lambda/db"
	"github.com/jinzhu/gorm"
)

type AdditionalItemsGroup struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255);not null;"`
	Description     string `gorm:"type:varchar(255);"`
	Order           int    `gorm:"type:int;not null;"`
	MinimumItems    int    `gorm:"type:int;not null;"`
	MaximumItems    int    `gorm:"type:int;not null;"`
	Enabled         bool   `gorm:"type:boolean;not null;"`
	CompanyID       uint   `gorm:"type:bigint;not null;"`
	AdditionalItems []*AdditionalItem
}

func (AdditionalItemsGroup) TableName() string {
	return "additional_items_group"
}

func (o AdditionalItemsGroup) AsJSON() *AdditionalItemsGroupJSON {
	additonalItemJSON := make([]*AdditionalItemJSON, 0)
	for _, additionalItem := range o.AdditionalItems {
		additonalItemJSON = append(additonalItemJSON, additionalItem.asJSON())
	}
	return &AdditionalItemsGroupJSON{o.ID, o.Name, o.Description, o.Order, o.MinimumItems, o.MaximumItems, o.Enabled, CompanyMinimalJSON{ID: o.CompanyID}, additonalItemJSON}
}

func (g *AdditionalItemsGroup) GetAssociationIDs() []uint {
	additionalItems := make([]uint, 0)
	for _, additionalItem := range g.AdditionalItems {
		additionalItems = append(additionalItems, additionalItem.ID)
	}
	return additionalItems
}

func (o AdditionalItemsGroup) GetItemsToRemove(firstSlice []uint, otherSlice []uint) []uint {
	notUsedMap := make(map[uint]bool)
	notUsedSlice := make([]uint, 0)
	for _, value := range firstSlice {
		found := false
		for _, otherValue := range otherSlice {
			if value == otherValue {
				found = true
			}
		}
		if !found {
			notUsedMap[value] = true
		}
	}
	for k := range notUsedMap {
		notUsedSlice = append(notUsedSlice, k)
	}
	return notUsedSlice
}

type AdditionalItemsGroupJSON struct {
	ID              uint                  `json:"ID"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	Order           int                   `json:"order"`
	MinimumItems    int                   `json:"minimumItems"`
	MaximumItems    int                   `json:"maximumItems"`
	Enabled         bool                  `json:"enabled"`
	Company         CompanyMinimalJSON    `json:"company"`
	AdditionalItems []*AdditionalItemJSON `json:"additionalItems"`
}

func (o AdditionalItemsGroupJSON) AsModel() *AdditionalItemsGroup {
	additionalItems := make([]*AdditionalItem, 0)
	for _, additionalItemJSON := range o.AdditionalItems {
		additionalItems = append(additionalItems, additionalItemJSON.AsModel())
	}
	return &AdditionalItemsGroup{gorm.Model{ID: o.ID}, o.Name, o.Description, o.Order, o.MinimumItems, o.MaximumItems, o.Enabled, o.Company.ID, additionalItems}
}

func (g AdditionalItemsGroupJSON) GetId() uint {
	return g.ID
}

func SaveAdditionalItemsGroup(groupToSave AdditionalItemsGroup) (AdditionalItemsGroup, error) {
	if !db.DB.NewRecord(&groupToSave) {
		return AdditionalItemsGroup{}, errors.New("Could not update a request with a POST")
	}
	if additionalItemsGroupExists(groupToSave) {
		return AdditionalItemsGroup{}, errors.New("Group with name " + groupToSave.Name + " already exists")
	}
	db.DB.Save(&groupToSave)
	return groupToSave, nil
}

func UpdateAdditionalItemsGroup(groupToUpdate AdditionalItemsGroup) (AdditionalItemsGroup, error) {

	if additionalItemsGroupExists(groupToUpdate) {
		return AdditionalItemsGroup{}, errors.New("Group with name " + groupToUpdate.Name + " already exists")
	}
	databaseGroup := AdditionalItemsGroup{}
	db.DB.Where("id = ? ", groupToUpdate.ID).Preload("AdditionalItems").Find(&databaseGroup)
	itemsToRemove := databaseGroup.GetItemsToRemove(databaseGroup.GetAssociationIDs(), groupToUpdate.GetAssociationIDs())
	db.DB.Unscoped().Where("id IN (?)", itemsToRemove).Delete(AdditionalItem{})
	db.DB.Save(&groupToUpdate)
	db.DB.Preload("AdditionalItems").Where("id = ?", groupToUpdate.ID).Find(&groupToUpdate)
	return groupToUpdate, nil
}

func ListAdditionalItemsByCompanyId(companyID string) []*AdditionalItemsGroupJSON {
	var groups []*AdditionalItemsGroup
	db.DB.Preload("AdditionalItems").Where("company_id = ?", companyID).Find(&groups)
	groupJSON := make([]*AdditionalItemsGroupJSON, 0)
	for _, group := range groups {
		groupJSON = append(groupJSON, group.AsJSON())
	}
	return groupJSON
}

func DeleteAdditionalGroup(groupId string) error {
	id, err := strconv.ParseInt(groupId, 10, 64)
	if err != nil {
		return err
	}
	db.DB.Unscoped().Where("id = ?", id).Delete(&AdditionalItemsGroup{})
	return nil
}

func additionalItemsGroupExists(group AdditionalItemsGroup) bool {
	var groupFromDb AdditionalItemsGroup
	db.DB.Raw("SELECT * FROM additional_items_group WHERE name = ? and id != ? and company_id = ?", group.Name, group.ID, group.CompanyID).Scan(&groupFromDb)
	return groupFromDb.ID != 0
}

func (t *AdditionalItemsGroup) HasConflict() (error, []string) {
	return t.hasSameItem()
}

func (t *AdditionalItemsGroup) additionalItemsGroupExists() bool {
	var loaded AdditionalItemsGroup
	db.DB.Where("name = ? and id != ? and company_id = ?", strings.TrimSpace(t.Name), t.ID, t.CompanyID).Find(&loaded)
	return loaded.ID != 0
}

func (t *AdditionalItemsGroup) hasSameItem() (error, []string) {
	if t.additionalItemsGroupExists() {
		return errors.New("additionalitemsgroup.errors.alreadyexists"), []string{t.Name}
	}
	return nil, nil
}
