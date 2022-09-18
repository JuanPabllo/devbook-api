package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		responses.Error(res, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(); erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositories := repositories.NewUserRepository(db)

	user.ID, erro = repositories.Create(user)
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(res, http.StatusCreated, user)
}

func SearchUsers(res http.ResponseWriter, req *http.Request) {
	nameOrNick := strings.ToLower(req.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositories := repositories.NewUserRepository(db)
	user, erro := repositories.Search(nameOrNick)
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(res, http.StatusOK, user)
}

func SearchUserById(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("search user by id"))
}

func UpdateUserbyId(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("update user by id"))
}

func DeleteUserById(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete user by id"))
}
