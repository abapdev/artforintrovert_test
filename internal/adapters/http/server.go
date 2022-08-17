package http

import (
	"artforintrovert_test/internal/domain/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *Rest) serviceRoutes(router chi.Router) {
	router.Get("/", r.List)
	router.Patch("/delete", r.Update)
	router.Delete("/delete", r.Delete)
}
func (r *Rest) List(writer http.ResponseWriter, request *http.Request) {

}
func (r *Rest) Update(writer http.ResponseWriter, request *http.Request) {

}

func (r *Rest) Delete(writer http.ResponseWriter, request *http.Request) {

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
