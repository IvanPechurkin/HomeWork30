package handlers

import (
	"30/pkg/api"
	"30/pkg/storage/interfaces"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetFriends(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Формирование ответа
		var status int
		var data []byte
		userId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/friends/"))

		users, err := repo.GetFriends(context.TODO(), userId)
		if err != nil {
			data, _ = json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal(api.ResponseDTO{
				Message: "Операция выполнена успешно",
				Items:   users,
			})
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
