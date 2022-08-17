package http

import (
	"artforintrovert_test/internal/domain/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *Rest) serviceRoutes(router chi.Router) {
	router.Get("/", r.List)
	router.Patch("/update", r.Update)
	router.Delete("/delete", r.Delete)
}
func (r *Rest) List(writer http.ResponseWriter, request *http.Request) {
	//вытаскиваем имя из строки запроса
	name := request.URL.Query().Get("name")
	//получаем от сервиса данные по запросу
	list, err := r.Service.List(context.Background(), name)
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	//Пишем ответ
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(list)
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	_, _ = writer.Write(jsonResp)

}
func (r *Rest) Update(writer http.ResponseWriter, request *http.Request) {
	//вытаскиваем имя из строки запроса
	name := request.URL.Query().Get("name")
	if name == "" {
		r.prepareErrResponse(writer, http.StatusBadRequest, errors.New("Empty name"))
		return
	}
	//вытаскиваем позицию из тела запроса
	var updateData *models.DataJSON
	if err := json.NewDecoder(request.Body).Decode(&updateData); err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	//сервис обновляет данные
	err := r.Service.Update(context.Background(), updateData)
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	//Пишем ответ
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	resp := models.DataState{
		Name:  updateData.Name,
		State: "Updated",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	_, _ = writer.Write(jsonResp)
}

func (r *Rest) Delete(writer http.ResponseWriter, request *http.Request) {
	//вытаскиваем имя из строки запроса
	name := request.URL.Query().Get("name")
	if name == "" {
		r.prepareErrResponse(writer, http.StatusBadRequest, errors.New("Empty name"))
		return
	}
	//сервис удаляет запись
	err := r.Service.Delete(context.Background(), &models.DataJSON{Name: name})
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	//Пишем ответ
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	resp := models.DataState{
		Name:  name,
		State: "Deleted",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		r.prepareErrResponse(writer, http.StatusBadRequest, err)
		return
	}
	_, _ = writer.Write(jsonResp)
}
func (r *Rest) prepareErrResponse(w http.ResponseWriter, httpCode int, message error) {
	logrus.WithError(message).Warn("Service internal error")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	jsonResp, err := json.Marshal(models.Error{
		Message: message.Error(),
	})
	if err != nil {
		logrus.Warn("Error happened in JSON response marshal")
	}
	_, _ = w.Write(jsonResp)
}
