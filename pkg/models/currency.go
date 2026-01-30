package models

import "time"

type Currency struct {
	ISOCode     int    `gorm:"primaryKey;not null"`
	ISOCharCode string `gorm:"size:3"`
	CbCode      string `gorm:"size:7"`
	Name        string `gorm:"size:100"`
	LatName     string `gorm:"size:100"`
	Nominal     int
	ParentCode  string `gorm:"size:7"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
