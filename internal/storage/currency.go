package storage

import (
	"fmt"

	"github.com/boldlogic/cbr-market-data-worker/internal/models"
)

func (st *Storage) saveCurrency(row *models.Currency) error {
	if row.ISOCode <= 0 {
		return fmt.Errorf("ISOCode должен быть больше 0. CbCode: %v, ISOCharCode: %v", row.CbCode, row.ISOCharCode)
	}

	result := st.db.Where(models.Currency{ISOCode: row.ISOCode}).Assign(models.Currency{
		ISOCharCode: row.ISOCharCode,
		CbCode:      row.CbCode,
		Name:        row.Name,
		LatName:     row.LatName,
		Nominal:     row.Nominal,
		ParentCode:  row.ParentCode,
		ISOCode:     row.ISOCode,
	}).FirstOrCreate(&row)
	if result.Error != nil {
		return fmt.Errorf("Currency. %v", result.Error)
	}

	return nil
}

func (st *Storage) SaveCurrencies(rows []models.Currency) []error {
	var errs []error
	for _, row := range rows {
		err := st.saveCurrency(&row)
		if err != nil {
			errs = append(errs, err)
		}

	}
	return errs

}
