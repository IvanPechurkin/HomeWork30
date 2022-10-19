package handlers

import (
	"30/pkg/api"
	"30/pkg/model"
	"30/pkg/storage/interfaces"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Add(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var u model.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		repo.AddUser(context.TODO(), &u)

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Пользователь создан",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
