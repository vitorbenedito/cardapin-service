package model

import (
	"errors"
	"net/url"

	"cardap.in/awsenvironment"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name                  string                  `gorm:"type:varchar(255);not null;"`
	Description           string                  `gorm:"type:varchar(255);"`
	Price                 int64                   `gorm:"type:bigint;"`
	Order                 int                     `gorm:"type:int;"`
	Enabled               bool                    `gorm:"type:boolean;"`
	Image                 string                  `gorm:"type:varchar(255);"`
	AdultsOnly            bool                    `gorm:"type:boolean;"`
	AdditionalItemsGroups []*AdditionalItemsGroup `gorm:"many2many:product_additional_group;"`
	CategoryID            uint                    `gorm:"type:bigint; not null;"`
}

type ProductJSON struct {
	ID                    uint                        `json:"id"`
	Name                  string                      `json:"name"`
	Description           string                      `json:"description"`
	Price                 int64                       `json:"price"`
	Order                 int                         `json:"order"`
	Enabled               bool                        `json:"enabled"`
	Image                 string                      `json:"image"`
	ImageURL              *string                     `json:"imageUrl"`
	AdultsOnly            bool                        `json:"adultsOnly"`
	AdditionalItemsGroups []*AdditionalItemsGroupJSON `json:"additionalItemsGroups"`
}

func (p *Product) HideDisabled() {
	groupsToShow := make([]*AdditionalItemsGroup, 0)
	for _, group := range p.AdditionalItemsGroups {
		if group.Enabled {
			additionalItemsToShow := make([]*AdditionalItem, 0)
			for _, item := range group.AdditionalItems {
				if item.Enabled {
					additionalItemsToShow = append(additionalItemsToShow, item)
				}
			}
			group.AdditionalItems = additionalItemsToShow
			groupsToShow = append(groupsToShow, group)
		}
	}
	p.AdditionalItemsGroups = groupsToShow
}

func (p *Product) hasSameAddionalItem() error {
	names := map[string]bool{}
	for _, a := range p.AdditionalItemsGroups {
		names[a.Name] = true
	}
	if len(p.AdditionalItemsGroups) != len(names) {
		return errors.New("additionalitem.errors.alreadyexists")
	}
	return nil
}

func (p *Product) GetGroupsIds() []*uint {
	groups := make([]*uint, 0)
	for _, group := range p.AdditionalItemsGroups {
		groups = append(groups, &group.ID)
	}
	return groups
}

func (p *Product) AsJSON() *ProductJSON {
	additonalItemGroupJSON := make([]*AdditionalItemsGroupJSON, 0)
	for _, additionalItem := range p.AdditionalItemsGroups {
		additonalItemGroupJSON = append(additonalItemGroupJSON, additionalItem.AsJSON())
	}
	var imageURLPointer *string
	if p.Image != "" {
		imageURL := awsenvironment.AwsUrl + "/" + url.QueryEscape(p.Image)
		imageURLPointer = &imageURL
	}
	return &ProductJSON{
		p.ID,
		p.Name,
		p.Description,
		p.Price,
		p.Order,
		p.Enabled,
		p.Image,
		imageURLPointer,
		p.AdultsOnly,
		additonalItemGroupJSON,
	}
}

func (pj *ProductJSON) AsModel(categoryId uint) *Product {
	additonalItemsGroup := make([]*AdditionalItemsGroup, 0)
	for _, additionalItemsGroupJSON := range pj.AdditionalItemsGroups {
		additonalItemsGroup = append(additonalItemsGroup, additionalItemsGroupJSON.AsModel())
	}
	return &Product{
		gorm.Model{ID: pj.ID},
		pj.Name,
		pj.Description,
		pj.Price,
		pj.Order,
		pj.Enabled,
		pj.Image,
		pj.AdultsOnly,
		additonalItemsGroup,
		categoryId,
	}
}

func (Product) TableName() string {
	return "product"
}

func (p ProductJSON) GetId() uint {
	return p.ID
}
