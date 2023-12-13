package services

import (
	"errors"
	"fmt"

	"cardap.in/auth"
	"cardap.in/db"
	"cardap.in/email"

	"cardap.in/apperrors"
	"cardap.in/model"
)

type CompanyServices struct {
}

func (c *CompanyServices) Save(companyToSave model.Company) (model.Company, error) {
	db.DB.Create(&companyToSave)

	if c.CompanyCodeExists(companyToSave) {
		return model.Company{}, errors.New("Company code with name " + companyToSave.CompanyCode + " already exists")
	}
	db.DB.Save(&companyToSave)
	return companyToSave, nil
}

func (c *CompanyServices) Update(companyToUpdate model.Company) (model.Company, error) {
	if c.CompanyCodeExists(companyToUpdate) {
		return model.Company{}, errors.New("Company code with name " + companyToUpdate.CompanyCode + " already exists")
	}
	companyFromDB := c.List(fmt.Sprint(companyToUpdate.ID))
	hours, adresses := companyToUpdate.GetAssociationIDs()
	hoursToRemove, adressesToRemove := companyFromDB.GetNotUsedIds(hours, adresses)
	db.DB.Unscoped().Where("id IN (?)", hoursToRemove).Delete(model.OpeningHours{})
	db.DB.Unscoped().Where("id IN (?)", adressesToRemove).Delete(model.Address{})
	var sections []*model.Section
	db.DB.Where("id IN (?)", companyToUpdate.GetSectionIds()).Find(&sections)
	companyToUpdate.Sections = sections
	var paymentType []*model.PaymentType
	db.DB.Where("id IN (?)", companyToUpdate.GetPaymentTypeIds()).Find(&paymentType)
	companyToUpdate.PaymentTypes = paymentType
	var additionalItemGroups []*model.AdditionalItemsGroup
	db.DB.Where("company_id = ?", companyToUpdate.ID).Find(&additionalItemGroups)
	companyToUpdate.AdditionalItemGroups = additionalItemGroups
	db.DB.Save(&companyToUpdate)
	return companyToUpdate, nil
}

func (*CompanyServices) List(id string) model.Company {
	company := model.Company{}
	db.DB.Preload("Sections").Preload("Addresses").Preload("PaymentTypes").Preload("OpeningHours").Where("ID = ?", id).Find(&company)
	return company
}

func (*CompanyServices) GetByIdFunction() func(id int) model.Company {
	getCompanyFunction := func(id int) model.Company {
		company := model.Company{}
		db.DB.Where("id = ?", id).Find(&company)
		return company
	}
	return getCompanyFunction
}

func (*CompanyServices) CompanyInterested(mailInfo email.Email) bool {
	sentToClient := email.Send(mailInfo.Email, true, email.ClientInterested, "Cardapin - Manifestação de Interesse", mailInfo)
	sentToCardapin := email.Send("oi@cardap.in", false, email.NewClient, "Novo Cliente - "+mailInfo.CompanyName, mailInfo)
	return sentToClient || sentToCardapin
}

func (*CompanyServices) GetCompanyByToken(token string) (*model.Company, *apperrors.AppError) {
	userJSON, err := auth.TokenValid(token)
	if err != nil {
		return nil, err
	}
	company := model.Company{}
	db.DB.Where("id = ?", userJSON.Company.ID).Find(&company)
	return &company, nil
}

func (*CompanyServices) CompanyCodeExists(company model.Company) bool {
	var companyFromDB model.Company
	db.DB.Raw("SELECT * FROM company WHERE company_code = ? and id != ?", company.CompanyCode, company.ID).Scan(&companyFromDB)
	return companyFromDB.ID != 0
}
