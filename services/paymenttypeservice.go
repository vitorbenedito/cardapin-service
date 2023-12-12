package services

import (
	"cardap.in/db"
	"cardap.in/model"
)

type PaymentTypeService struct {
}

func (pt *PaymentTypeService) ListPaymentTypes() ([]*model.PaymentTypeJSON, error) {
	var paymentTypes []*model.PaymentType
	db.DB.Find(&paymentTypes)
	paymentTypesJSON := make([]*model.PaymentTypeJSON, 0)
	for _, paymentType := range paymentTypes {
		paymentTypesJSON = append(paymentTypesJSON, paymentType.AsJSON())
	}
	return paymentTypesJSON, nil
}
