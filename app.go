package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	. "db_connect/dao"
	. "db_connect/config"
	. "db_connect/model"

)

var config = Config{}
var dao = MyDao{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
	fmt.Println("Connected")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/users", InsertUser).Methods("POST")

	if err:=http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	
	//defer r.Body.Close()
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.InsertUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}