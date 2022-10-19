package handlers

import (
	"30/pkg/api"
	"30/pkg/storage/interfaces"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Update(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data, _ := json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			w.Write(data)
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var t api.RequestDTO
		if err := json.Unmarshal(content, &t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data, _ := json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			w.Write(data)
			return
		}

		//Формирование ответа
		var status int
		var data []byte
		userId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			data, _ := json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			w.Write(data)
			return
		}
		if err := repo.UpdateUserAge(context.TODO(), userId, t.NewAge); err != nil {
			data, _ = json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal(api.ResponseDTO{
				Message: "Возраст пользователя обновлен",
			})
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
