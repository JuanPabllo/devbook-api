package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Prepare("register"); erro != nil {
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
	params := mux.Vars(req)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	user, erro := repositorie.SearchById(userID)
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(res, http.StatusOK, user)
}

func UpdateUserbyId(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

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

	if erro = user.Prepare("update"); erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	if erro = repositorie.Update(userID, user); erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(res, http.StatusNoContent, nil)
}

func DeleteUserById(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(res, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	if erro = repositorie.Delete(userID); erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(res, http.StatusNoContent, nil)
}
