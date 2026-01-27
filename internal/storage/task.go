package storage

import (
	"fmt"

	"github.com/boldlogic/cbr-market-data-worker/internal/models"
)

func (st *Storage) GetAction(code string) (models.Action, error) {
	var result models.Action

	err := st.db.Where("code=?", code).First(&result).Error

	if err != nil {

		return models.Action{}, err
	}
	return result, nil
}

func (st *Storage) SaveAction(row *models.Action) error {

	result := st.db.Where(models.Action{Code: row.Code}).Assign(models.Action{
		Id:   row.Id,
		Code: row.Code,
		Name: row.Name,
	}).FirstOrCreate(row)
	if result.Error != nil {
		return fmt.Errorf("Произошла ошибка при сохранении Action. %w", result.Error)
	}

	return nil
}

// func (st *Storage) SaveTask(row *models.Task) error {

// 	result := st.db.Where(models.Task{: row.Code}).Assign(models.Task{
// 		Id:   row.Id,
// 		Code: row.Code,
// 		Name: row.Name,
// 	}).FirstOrCreate(row)
// 	if result.Error != nil {
// 		return fmt.Errorf("Произошла ошибка при сохранении Action. %w", result.Error)
// 	}

// 	return nil
// }

func (st *Storage) GetTask(uuid string) (models.Task, error) {
	var result models.Task

	err := st.db.Where("uuid=?", uuid).First(&result).Error

	if err != nil {

		return models.Task{}, err
	}
	return result, nil
}

const (
	CreateTask = `
	insert into tasks (uuid, actionId, char_code, date_from, date_to)
	values (?, ?,?,?,?)
	`
)

func (st *Storage) CreateTask(row *models.Task) error {

	result := st.db.Exec(CreateTask, row.Uuid, row.ActionId, row.CharCode, row.DateFrom, row.DateTo)

	if result.Error != nil {
		return fmt.Errorf("Произошла ошибка при сохранении Task. %w", result.Error)
	}

	return nil
}
