package model

import "github.com/jinzhu/gorm"

type Section struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(255); not null;"`
	Companies []*Company `gorm:"many2many:company_section;"`
}

type SectionJSON struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s Section) AsJSON() *SectionJSON {
	return &SectionJSON{s.ID, s.Name}
}

func (s SectionJSON) AsModel() *Section {
	return &Section{gorm.Model{ID: s.ID},
		s.Name, make([]*Company, 0)}
}

func (Section) TableName() string {
	return "section"
}
