package ports

import (
	"net/http"
)

type ServiceHTTP interface {
	//обработчик эндпоинта GET list
	List(writer http.ResponseWriter, request *http.Request)
	//обработчик эндпоинта Patch name
	Update(writer http.ResponseWriter, request *http.Request)
	//обработчик эндпоинта Delete name
	Delete(writer http.ResponseWriter, request *http.Request)
}
