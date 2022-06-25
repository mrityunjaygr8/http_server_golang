package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mrtyunjaygr8/http_server_golang/utils"
)

type createPostParams struct {
	UserEmail string `json:"userEmail"`
	Text      string `json:"text"`
}

func (a ApiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	params := createPostParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	post, err := a.DbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, post)
}

func (a ApiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userEmail := strings.TrimPrefix(path, "/posts/")

	posts, err := a.DbClient.GetPosts(userEmail)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}

func (a ApiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/posts/")

	err := a.DbClient.DeletePost(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, struct{}{})
}
