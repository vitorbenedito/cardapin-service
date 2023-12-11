package model

type Client struct {
	Phone      uint64 `gorm:"primary_key; not_null; auto_increment:false;type:bigint;" json:"phone"`
	Name       string `gorm:"type:varchar(255);not null;" json:"name"`
	Address    string `gorm:"type:varchar(255);" json:"address"`
	Area       string `gorm:"type:varchar(255);" json:"area"`
	City       string `gorm:"type:varchar(255);" json:"city"`
	State      string `gorm:"type:varchar(255);" json:"state"`
	PostalCode string `gorm:"type:varchar(255);" json:"postalCode"`
	Complement string `gorm:"type:varchar(255);" json:"complement"`
	Number     string `gorm:"type:varchar(10);" json:"number"`
	Landmark   string `gorm:"type:varchar(255);" json:"landmark"`
}

func (c Client) GetId() uint {
	return uint(c.Phone)
}
