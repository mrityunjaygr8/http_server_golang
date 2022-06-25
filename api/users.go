package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/mrtyunjaygr8/http_server_golang/utils"
)

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return errors.New("email can't be empty")
	}

	if password == "" {
		return errors.New("password can't be empty")
	}

	if age < 18 {
		return errors.New("age must be at least 18 years")
	}

	return nil

}

type createUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

func (a ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params := createUserParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = userIsEligible(params.Email, params.Password, params.Age)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := a.DbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

type updateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

func (a ApiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	params := createUserParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = userIsEligible(params.Email, params.Password, params.Age)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	user, err := a.DbClient.UpdateUser(params.Email, params.Password, params.Name, params.Age)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
func (a ApiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	email := strings.TrimPrefix(path, "/users/")
	err := a.DbClient.DeleteUser(email)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err)
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}

func (a ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	email := strings.TrimPrefix(path, "/users/")
	user, err := a.DbClient.GetUser(email)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err)
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}
