package service

import (
	"context"
	"fmt"

	"github.com/boldlogic/cbr-market-data-worker/internal/models"
	UUID "github.com/google/uuid"
)

func (c *Service) CreateTask(ctx context.Context, task models.Task, action string) error {
	act, err := c.schedulerRepo.GetAction(action)
	if err != nil {
		return err
	}
	if act.Id == 0 {
		return fmt.Errorf("Не найден тип события с кодом %s", action)
	}
	task.ActionId = act.Id

	if task.Uuid != "" {
		exists, err := c.schedulerRepo.GetTask(task.Uuid)
		if err != nil {
			return fmt.Errorf("Не удалось получить задачу по uuid %s", task.Uuid)
		}
		if exists.Id != 0 {
			return fmt.Errorf("Задача с uuid %s уже существует, id=%d", task.Uuid, exists.Id)
		}
	} else {
		task.Uuid = UUID.New().String()
		err = c.schedulerRepo.CreateTask(&task)
		if err != nil {
			return fmt.Errorf("Не удалось создать задачу по uuid %s", task.Uuid)
		}
	}

	return nil
}
