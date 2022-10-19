package handlers

import (
	"30/pkg/api"
	"30/pkg/storage/interfaces"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Link(repo interfaces.Repository) http.HandlerFunc {
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
		if err := repo.LinkUsers(context.TODO(), t.Source, t.Target); err != nil {
			data, _ = json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal(api.ResponseDTO{
				Message: "Пользователи добавлены в друзья друг к другу",
			})
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
