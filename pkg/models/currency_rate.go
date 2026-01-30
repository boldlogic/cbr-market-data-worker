package models

import "time"

type FxRate struct {
	Date             time.Time `gorm:"type:date;primaryKey;not null"`
	QuoteISOCode     int       `gorm:"primaryKey;not null"`
	BaseISOCode      int       `gorm:"primaryKey;not null"`
	Nominal          int
	QuoteForNominal  float64
	QuotePerUnit     float64
	BasePerQuoteUnit float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
