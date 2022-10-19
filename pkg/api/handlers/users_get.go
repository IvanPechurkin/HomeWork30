package handlers

import (
	"30/pkg/api"
	"30/pkg/storage/interfaces"
	"context"
	"encoding/json"
	"net/http"
)

func GetAll(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repo.GetAll(context.TODO())

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Операция выполнена успешно",
			Items:   users,
		})
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
