package model

import (
	"time"

	"gorm.io/gorm"
)

type OpeningHours struct {
	gorm.Model
	StartDay  string    `gorm:"type:varchar(255);not null;"`
	EndDay    string    `gorm:"type:varchar(255);not null;"`
	StartTime time.Time `gorm:"type:time(4);not null;"`
	EndTime   time.Time `gorm:"type:time(4);not null;"`
	CompanyID uint      `gorm:"type:bigint;not null;"`
}

type OpeningHoursJSON struct {
	ID        uint   `json:"id"`
	StartDay  string `json:"startDay"`
	EndDay    string `json:"endDay"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (s OpeningHours) AsJSON() *OpeningHoursJSON {
	return &OpeningHoursJSON{s.ID, s.StartDay, s.EndDay, s.StartTime.Format("15:04"), s.EndTime.Format("15:04")}
}

func (s OpeningHoursJSON) AsModel() *OpeningHours {
	startTime, _ := time.Parse("15:04", s.StartTime)
	endTime, _ := time.Parse("15:04", s.EndTime)
	return &OpeningHours{gorm.Model{ID: s.ID},
		s.StartDay, s.EndDay, startTime, endTime, 0}
}

func (OpeningHours) TableName() string {
	return "opening_hours"
}
