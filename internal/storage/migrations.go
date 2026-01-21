package storage

import (
	"fmt"
	"strconv"

	"github.com/bojanz/currency"
	"github.com/boldlogic/cbr-market-data-worker/internal/models"
)

func (st *Storage) Migrate() error {
	err := st.db.AutoMigrate(&models.Currency{})
	if err != nil {
		return err
	}
	err = st.db.AutoMigrate(&models.FxRate{})
	if err != nil {
		return err
	}
	err = st.FillRussianRuble()
	if err != nil {
		return err
	}
	return nil

}

func (st *Storage) FillRussianRuble() error {
	rubCode := "RUB"
	numCodeStr, ok := currency.GetNumericCode(rubCode)
	if !ok {
		return fmt.Errorf("Не удалось получить код валюты для RUB из библиотеки")
	}

	numCode, err := strconv.Atoi(numCodeStr)
	if err != nil {
		return fmt.Errorf("не удалось сконвертировать код валюты: %w", err)
	}

	rub := models.Currency{
		ISOCharCode: rubCode,
		CbCode:      "",
		Name:        "Российский рубль",
		LatName:     "Russian Ruble",
		Nominal:     1,
		ParentCode:  "",
		ISOCode:     numCode,
	}
	result := st.db.Where(models.Currency{ISOCode: numCode}).Assign(rub).FirstOrCreate(&rub)
	if result.Error != nil {
		return fmt.Errorf("%w", result.Error)
	}

	return nil
}
