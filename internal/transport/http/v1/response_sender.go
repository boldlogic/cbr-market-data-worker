package v1

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       interface{}
}

type Body struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (h *Handler) SendResponse(w http.ResponseWriter, response APIResponse) {

	if len(response.Headers) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		for k, v := range response.Headers {
			w.Header().Set(k, v)
		}
	}
	w.WriteHeader(response.StatusCode)

	if response.Body != nil {
		body, err := json.Marshal(response.Body)

		if response.StatusCode >= 500 {
			h.log.WithFields(logrus.Fields{
				"StatusCode": response.StatusCode,
				"Body":       response.Body,
			}).Error("Ответ с ошибкой (ошибка сервера)")
		} else if response.StatusCode >= 400 {
			h.log.WithFields(logrus.Fields{
				"StatusCode": response.StatusCode,
				"Body":       response.Body,
			}).Warn("Ответ с ошибкой (ошибка клиента)")
		}

		_, _ = w.Write(body)
		if err != nil {
			http.Error(w, "Произошла непредвиденная ошибка", http.StatusInternalServerError)
			h.log.Errorf("Не удалось сериализовать тело ответа: %v", err)
		}
	}
}
