package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/boldlogic/cbr-market-data-worker/internal/models"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Получен запрос на выполнение операции CreateTask")
	ctx := r.Context()
	reqTask, action, err := parseTask(r.Body)

	if err != nil {
		h.log.Errorf("Ошибка: %w", err)
		h.SendResponse(w, APIResponse{
			StatusCode: http.StatusBadRequest,
			Body: Body{
				Error: "Не удалось прочитать тело запроса",
			},
		})

		return
	}
	err = h.Service.CreateTask(ctx, reqTask, action)
	if err != nil {
		h.SendResponse(w, APIResponse{
			StatusCode: http.StatusConflict,
			Body: Body{
				Error: "Задача уже существует",
			},
		})

		return
	}
	h.SendResponse(w, APIResponse{
		StatusCode: http.StatusAccepted,
	})
}

func parseTask(r io.ReadCloser) (models.Task, string, error) {
	reqTask, err := getTaskFromBody(r)
	if err != nil {
		return models.Task{}, "", err
	}
	err = validateNewTask(reqTask)
	if err != nil {
		return models.Task{}, "", err
	}

	action := reqTask.Action

	var dateFrom *time.Time
	var dateTo *time.Time
	if reqTask.Params.DateFrom != "" {

		parsed, err := time.Parse(models.DateFormat, reqTask.Params.DateFrom) //2006-01-11
		if err != nil {
			return models.Task{}, "", err
		}
		dateFrom = &parsed
	}
	if reqTask.Params.DateTo != "" {

		parsed, err := time.Parse(models.DateFormat, reqTask.Params.DateTo) //2006-01-11
		if err != nil {
			return models.Task{}, "", err
		}
		dateTo = &parsed
	}
	task := models.Task{
		Uuid: reqTask.Uuid,
		//ActionId: models.Action(reqTask.Action),
		CharCode: &reqTask.Params.CcyCode,
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}
	return task, action, nil
}

func getTaskFromBody(r io.ReadCloser) (TaskDTO, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r)
	if err != nil {
		return TaskDTO{}, fmt.Errorf("Не удалось прочитать тело запроса: %v", err)
	}
	var tsk TaskDTO
	err = json.Unmarshal(buf.Bytes(), &tsk)
	if err != nil {
		return TaskDTO{}, fmt.Errorf("Не удалось распарсить тело запроса: %v", err)
	}
	return tsk, nil
}

func validateNewTask(task TaskDTO) error {
	if task.Action == "" {
		return fmt.Errorf("Поле action обязательно")
	}
	return nil
}
