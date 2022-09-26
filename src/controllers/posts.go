package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePosts(req http.ResponseWriter, res *http.Request) {
	userID, erro := auth.ExtractUserID(res)
	if erro != nil {
		responses.Error(req, http.StatusUnauthorized, erro)
		return
	}

	requestBody, erro := ioutil.ReadAll(res.Body)
	if erro != nil {
		responses.Error(req, http.StatusUnprocessableEntity, erro)
		return
	}

	var posts models.Posts
	if erro = json.Unmarshal(requestBody, &posts); erro != nil {
		responses.Error(req, http.StatusBadRequest, erro)
		return
	}

	posts.AuthorID = userID

	if erro = posts.Preparar(); erro != nil {
		responses.Error(req, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorie := repositories.NewPostsRepository(db)
	posts.ID, erro = repositorie.Create(posts)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusCreated, posts)
}

func GetPosts(req http.ResponseWriter, res *http.Request) {}

func GetPostByID(req http.ResponseWriter, res *http.Request) {}

func UpdatePost(req http.ResponseWriter, res *http.Request) {}

func DeletePost(req http.ResponseWriter, res *http.Request) {}
