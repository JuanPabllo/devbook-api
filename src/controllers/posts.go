package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetPosts(req http.ResponseWriter, res *http.Request) {
	userID, erro := auth.ExtractUserID(res)
	if erro != nil {
		responses.Error(req, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorie := repositories.NewPostsRepository(db)
	posts, erro := repositorie.SearchPosts(userID)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusOK, posts)
}

func GetPostByID(req http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)
	postID, erro := strconv.ParseUint(params["postID"], 10, 64)
	if erro != nil {
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
	post, erro := repositorie.GetPostByID(postID)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusOK, post)
}

func UpdatePost(req http.ResponseWriter, res *http.Request) {
	userID, erro := auth.ExtractUserID(res)
	if erro != nil {
		responses.Error(req, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(res)
	postID, erro := strconv.ParseUint(params["postID"], 10, 64)
	if erro != nil {
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
	postSaveInDB, erro := repositorie.GetPostByID(postID)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	if postSaveInDB.AuthorID != userID {
		responses.Error(req, http.StatusForbidden, errors.New("you can not edit this post"))
		return
	}

	requestBody, erro := ioutil.ReadAll(res.Body)
	if erro != nil {
		responses.Error(req, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Posts

	if erro = json.Unmarshal(requestBody, &post); erro != nil {
		responses.Error(req, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Preparar(); erro != nil {
		responses.Error(req, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorie.Update(postID, post); erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusNoContent, nil)
}

func DeletePost(req http.ResponseWriter, res *http.Request) {
	userID, erro := auth.ExtractUserID(res)
	if erro != nil {
		responses.Error(req, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(res)
	postID, erro := strconv.ParseUint(params["postID"], 10, 64)
	if erro != nil {
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
	postSaveInDB, erro := repositorie.GetPostByID(postID)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	if postSaveInDB.AuthorID != userID {
		responses.Error(req, http.StatusForbidden, errors.New("you can not delete this post"))
		return
	}

	if erro = repositorie.Delete(postID); erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusNoContent, nil)
}

func GetPostsByUser(req http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)
	userID, erro := strconv.ParseUint(params["userID"], 10, 64)
	if erro != nil {
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
	posts, erro := repositorie.GetPostsByUser(userID)
	if erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusOK, posts)
}

func LikePost(req http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)
	postID, erro := strconv.ParseUint(params["postID"], 10, 64)
	if erro != nil {
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
	if erro := repositorie.Like(postID); erro != nil {
		responses.Error(req, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(req, http.StatusNoContent, nil)
}
