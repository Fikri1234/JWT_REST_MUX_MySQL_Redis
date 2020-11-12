package service

import (
	"JWT_REST_MySQL_JWT_Redis/model"
	"JWT_REST_MySQL_JWT_Redis/repository"
	"JWT_REST_MySQL_JWT_Redis/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

// UserDetailRoutes route
var UserDetailRoutes = model.RoutePrefix{
	Prefix: "/api",
	SubRoutes: []model.Route{
		model.Route{
			Name:        "GetUserDetailByID",
			Method:      "GET",
			Pattern:     "/user/dtl/{id}",
			HandlerFunc: GetUserDetailByID,
			Protected:   true,
		},
		model.Route{
			Name:        "GetAllUserDetails",
			Method:      "GET",
			Pattern:     "/user/dtl/",
			HandlerFunc: GetAllUserDetails,
			Protected:   true,
		},
		model.Route{
			Name:        "CreateUserDetail",
			Method:      "POST",
			Pattern:     "/user/dtl/",
			HandlerFunc: CreateUserDetail,
			Protected:   true,
		},
		model.Route{
			Name:        "UpdateUserDetail",
			Method:      "PUT",
			Pattern:     "/user/dtl/",
			HandlerFunc: UpdateUserDetail,
			Protected:   true,
		},
		model.Route{
			Name:        "DeleteUserDetailByID",
			Method:      "DELETE",
			Pattern:     "/user/dtl/{id}",
			HandlerFunc: DeleteUserDetailByID,
			Protected:   true,
		},
	},
}
var pool *redis.Pool

// GetUserDetailByID by id
func GetUserDetailByID(w http.ResponseWriter, r *http.Request) {

	pool = &redis.Pool{
		MaxActive: 50,
		MaxIdle:   1000,
		// IdleTimeout: 240*time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "uutes", "123")
	if err != nil {
		panic(err)
	}

	redisIDUser, err := redis.String(conn.Do("GET", "uutes"))
	fmt.Printf("red uuid: %v", redisIDUser)
	if err != nil {
		panic(err)
	}

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Printf("lennn: %v", len(keys))
	for _, key := range keys {
		fmt.Println(key)
	}

	_, err = conn.Do("DEL", "uutes")
	if err != nil {
		panic(err)
	}

	var userDtl model.MUserDetail
	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userDtl, err = repository.GetUserDetailByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if (model.MUserDetail{}) == userDtl {
		util.ResponseWithJSON(w, http.StatusNotFound, userDtl)
	} else {
		util.ResponseWithJSON(w, http.StatusOK, userDtl)
	}
}

// GetAllUserDetails data
func GetAllUserDetails(w http.ResponseWriter, r *http.Request) {

	var userDtls []model.MUserDetail
	userDtls, err := repository.GetAllUserDetail()
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, userDtls)
}

// CreateUserDetail post
func CreateUserDetail(w http.ResponseWriter, r *http.Request) {

	var userDtl model.MUserDetail

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &userDtl)

	userDtl, err = repository.CreateUserDetail(userDtl)
	if err != nil {
		util.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusCreated, userDtl)
}

// UpdateUserDetail update user
func UpdateUserDetail(w http.ResponseWriter, r *http.Request) {

	var userDtl model.MUserDetail

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &userDtl)

	usrDtl, err := repository.UpdateUserDetail(userDtl)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, usrDtl)
}

// DeleteUserDetailByID ...
func DeleteUserDetailByID(w http.ResponseWriter, r *http.Request) {

	var userDtl model.MUserDetail

	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	err = repository.DeleteUserDetailByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusNoContent, userDtl)
}
