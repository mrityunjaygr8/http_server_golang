package api

import (
	"errors"
	"net/http"

	"github.com/mrtyunjaygr8/http_server_golang/internal/database"
	"github.com/mrtyunjaygr8/http_server_golang/utils"
)

type ApiConfig struct {
	DbClient database.Client
}

func (a ApiConfig) EndpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.handlerGetUser(w, r)
	case http.MethodPost:
		a.handlerCreateUser(w, r)
	case http.MethodPut:
		a.handlerUpdateUser(w, r)
	case http.MethodDelete:
		a.handlerDeleteUser(w, r)
	default:
		utils.RespondWithError(w, 404, errors.New("method is not supported"))
	}
}
