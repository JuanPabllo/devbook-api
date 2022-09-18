package controllers

import "net/http"

func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Create user"))
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
