package services

import (
	"cardap.in/db"
	"cardap.in/model"
)

type ClientService struct {
}

func (c *ClientService) Save(clientToSave model.Client, isUpdate bool) (*model.Client, error) {
	db.DB.Save(&clientToSave)
	return &clientToSave, nil
}

func (*ClientService) GetByPhone(phone uint64) (model.Client, error) {
	var client model.Client
	db.DB.Where("phone = ?", phone).Find(&client)
	return client, nil
}
