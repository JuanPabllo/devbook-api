package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repositories := repositories.NewUserRepository(db)
	repositories.Create(user)
}

func SearchUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("search all users"))
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
