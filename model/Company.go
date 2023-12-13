package model

import (
	"net/url"

	"cardap.in/awsenvironment"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name                 string     `gorm:"type:varchar(255);not null;"`
	Description          string     `gorm:"type:varchar(255);"`
	ImagePath            string     `gorm:"type:varchar(255);"`
	PhoneNumber          string     `gorm:"type:varchar(255);"`
	CompanyCode          string     `gorm:"type:varchar(255);"`
	WhatsApp             string     `gorm:"type:varchar(255);"`
	HasDelivery          bool       `gorm:"type:boolean;"`
	HasWithdrawn         bool       `gorm:"type:boolean;"`
	Theme                string     `gorm:"type:varchar(255);"`
	Sections             []*Section `gorm:"many2many:company_section;"`
	Addresses            []*Address
	PaymentTypes         []*PaymentType `gorm:"many2many:company_payment_type;"`
	OpeningHours         []*OpeningHours
	AdditionalItemGroups []*AdditionalItemsGroup
	UserID               uint  `gorm:"type:bigint;"`
	User                 *User `gorm:"foreignkey:UserID"`
}

func (Company) TableName() string {
	return "company"
}

func (c *Company) GetSectionIds() []*uint {
	sections := make([]*uint, 0)
	for _, section := range c.Sections {
		sections = append(sections, &section.ID)
	}
	return sections
}

func (c *Company) GetPaymentTypeIds() []*uint {
	paymentTypes := make([]*uint, 0)
	for _, paymentType := range c.PaymentTypes {
		paymentTypes = append(paymentTypes, &paymentType.ID)
	}
	return paymentTypes
}

func (c *Company) GetAdditionalItemsGroupIds() []*uint {
	groups := make([]*uint, 0)
	for _, group := range c.AdditionalItemGroups {
		groups = append(groups, &group.ID)
	}
	return groups
}

type CompanyJson struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	CompanyCode  string              `json:"companyCode"`
	WhatsApp     string              `json:"whatsApp"`
	HasDelivery  bool                `json:"hasDelivery"`
	HasWithdrawn bool                `json:"hasWithdrawn"`
	Theme        string              `json:"theme"`
	ImagePath    string              `json:"image"`
	ImageURL     *string             `json:"imageUrl"`
	PhoneNumber  string              `json:"phone"`
	Sections     []*SectionJSON      `json:"sections"`
	Addresses    []*AddressJSON      `json:"addresses"`
	PaymentTypes []*PaymentTypeJSON  `json:"paymentTypes"`
	OpeningHours []*OpeningHoursJSON `json:"openingHours"`
}

type CompanyMinimalJSON struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CompanyCode string `json:"companyCode"`
}

func (cj *CompanyJson) AsModel() *Company {
	sections := make([]*Section, 0)
	for _, sectionJSON := range cj.Sections {
		sections = append(sections, sectionJSON.AsModel())
	}
	addresses := make([]*Address, 0)
	for _, addressJSON := range cj.Addresses {
		address := addressJSON.AsModel()
		address.CompanyID = cj.ID
		addresses = append(addresses, address)
	}
	paymentTypes := make([]*PaymentType, 0)
	for _, paymentTypeJSON := range cj.PaymentTypes {
		paymentTypes = append(paymentTypes, paymentTypeJSON.AsModel())
	}
	openingHoursSlice := make([]*OpeningHours, 0)
	for _, openingHoursJSON := range cj.OpeningHours {
		openingHours := openingHoursJSON.AsModel()
		openingHours.CompanyID = cj.ID
		openingHoursSlice = append(openingHoursSlice, openingHours)
	}
	c := &Company{}
	c.ID = cj.ID
	c.Name = cj.Name
	c.Description = cj.Description
	c.ImagePath = cj.ImagePath
	c.PhoneNumber = cj.PhoneNumber
	c.Sections = sections
	c.Addresses = addresses
	c.PaymentTypes = paymentTypes
	c.OpeningHours = openingHoursSlice
	c.CompanyCode = cj.CompanyCode
	c.HasDelivery = cj.HasDelivery
	c.HasWithdrawn = cj.HasWithdrawn
	c.WhatsApp = cj.WhatsApp
	c.Theme = cj.Theme
	return c
}

func (c *Company) AsJson() *CompanyJson {
	sectionsJSON := make([]*SectionJSON, 0)
	for _, section := range c.Sections {
		sectionsJSON = append(sectionsJSON, section.AsJSON())
	}
	addressesJSON := make([]*AddressJSON, 0)
	for _, address := range c.Addresses {
		addressesJSON = append(addressesJSON, address.asJSON())
	}
	paymentTypesJSON := make([]*PaymentTypeJSON, 0)
	for _, paymentTypeJSON := range c.PaymentTypes {
		paymentTypesJSON = append(paymentTypesJSON, paymentTypeJSON.AsJSON())
	}
	openingHoursSliceJSON := make([]*OpeningHoursJSON, 0)
	for _, openingHours := range c.OpeningHours {
		openingHoursSliceJSON = append(openingHoursSliceJSON, openingHours.AsJSON())
	}
	var imageURLPointer *string
	if c.ImagePath != "" {
		imageURL := awsenvironment.AwsUrl + "/" + url.QueryEscape(c.ImagePath)
		imageURLPointer = &imageURL
	}
	cj := &CompanyJson{}
	cj.ID = c.ID
	cj.Name = c.Name
	cj.Description = c.Description
	cj.ImagePath = c.ImagePath
	cj.ImageURL = imageURLPointer
	cj.PhoneNumber = c.PhoneNumber
	cj.Sections = sectionsJSON
	cj.Addresses = addressesJSON
	cj.PaymentTypes = paymentTypesJSON
	cj.OpeningHours = openingHoursSliceJSON
	cj.CompanyCode = c.CompanyCode
	cj.HasDelivery = c.HasDelivery
	cj.HasWithdrawn = c.HasWithdrawn
	cj.WhatsApp = c.WhatsApp
	cj.Theme = c.Theme
	return cj
}

func (c *Company) AsMinimalJSON() *CompanyMinimalJSON {
	return &CompanyMinimalJSON{
		ID:          c.ID,
		Name:        c.Name,
		CompanyCode: c.CompanyCode,
	}
}

func (c CompanyJson) GetId() uint {
	return c.ID
}

func (c *Company) GetAssociationIDs() ([]uint, []uint) {
	openingHours := make([]uint, 0)
	addresses := make([]uint, 0)
	for _, openingHour := range c.OpeningHours {
		openingHours = append(openingHours, openingHour.ID)
	}
	for _, address := range c.Addresses {
		addresses = append(addresses, address.ID)
	}
	return openingHours, addresses
}

func (c *Company) GetNotUsedIds(otherHours []uint, otherAdresses []uint) ([]uint, []uint) {
	openingHours, addresses := c.GetAssociationIDs()
	return getNotFoundValues(openingHours, otherHours), getNotFoundValues(addresses, otherAdresses)
}
