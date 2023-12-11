package services

import (
	"cardap.in/lambda/db"
	"cardap.in/lambda/model"
)

type SectionService struct {
}

func (pt *SectionService) ListSection() ([]*model.SectionJSON, error) {
	var Sections []*model.Section
	db.DB.Find(&Sections)
	sectionsJSON := make([]*model.SectionJSON, 0)
	for _, section := range Sections {
		sectionsJSON = append(sectionsJSON, section.AsJSON())
	}
	return sectionsJSON, nil
}
