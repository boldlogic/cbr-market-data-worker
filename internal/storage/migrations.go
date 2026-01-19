package storage

import "github.com/boldlogic/cbr-market-data-worker/internal/models"

func (st *Storage) Migrate() error {
	err := st.db.AutoMigrate(&models.Currency{})
	if err != nil {
		return err
	}
	return nil

}
