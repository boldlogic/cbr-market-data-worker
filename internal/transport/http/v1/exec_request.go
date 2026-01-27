package v1

// func (h *Handler) execRequest(w http.ResponseWriter, r *http.Request) {
// 	h.log.Info("Получен запрос на выполнение операции")
// 	ctx := r.Context()
// 	var buf bytes.Buffer

// 	_, err := buf.ReadFrom(r.Body)
// 	if err != nil {
// 		h.log.Warnf("Не удалось прочитать тело запроса: %v", err)
// 		h.SendResponse(w, APIResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Body: Body{
// 				Error: err.Error(),
// 			},
// 		})
// 		return
// 	}
// 	var cb CbRequest
// 	err = json.Unmarshal(buf.Bytes(), &cb)
// 	if err != nil {
// 		h.log.Warnf("Не удалось распарсить тело запроса: %v", err)
// 		h.SendResponse(w, APIResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Body: Body{
// 				Error: err.Error(),
// 			},
// 		})
// 		return
// 	}
// 	var params = make(map[string]string)
// 	if cb.CharCode != "" {
// 		params["cbcode"] = cb.CharCode
// 	}
// 	if cb.DateFrom != "" {
// 		params["datefrom"] = cb.DateFrom
// 	}
// 	if cb.DateTo != "" {
// 		params["dateto"] = cb.DateTo
// 	}
// 	tsk := models.Task{
// 		//Type:     cb.Type,
// 		//Params:   params,
// 		CharCode: cb.CharCode,
// 		//DateFrom: cb.DateFrom,
// 		//DateTo:   cb.DateTo,
// 		Uuid:     cb.Uuid,
// 	}
// 	err = h.Service.ExecuteTask(ctx, tsk)
// 	//h.log.Info(resp)

// 	h.SendResponse(w, APIResponse{
// 		StatusCode: http.StatusOK,
// 	})

// }
