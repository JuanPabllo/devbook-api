package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		responses.Error(res, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
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
	userSaveInDB, erro := repositorie.SearchByEmail(user.Email)
	if erro != nil {
		responses.Error(res, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(userSaveInDB.Password, user.Password); erro != nil {
		responses.Error(res, http.StatusUnauthorized, erro)
		return
	}

	token, _ := auth.CreateToken(userSaveInDB.ID)
	res.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}
