package service

import (
	"JWT_REST_MySQL_JWT_Redis/model"
	"JWT_REST_MySQL_JWT_Redis/repository"
	"JWT_REST_MySQL_JWT_Redis/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// CredentialsLogin Create a struct to read the username and password from the request body
type CredentialsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRoutes route
var UserRoutes = model.RoutePrefix{
	Prefix: "/api",
	SubRoutes: []model.Route{
		model.Route{
			Name:        "FindUserById",
			Method:      "GET",
			Pattern:     "/user/{id}",
			HandlerFunc: GetUserByID,
			Protected:   true,
		},
		model.Route{
			Name:        "GetUsers",
			Method:      "GET",
			Pattern:     "/user/",
			HandlerFunc: GetUsers,
			Protected:   true,
		},
		model.Route{
			Name:        "CreateUser",
			Method:      "POST",
			Pattern:     "/user/",
			HandlerFunc: CreateUser,
			Protected:   true,
		},
		model.Route{
			Name:        "UpdateUser",
			Method:      "PUT",
			Pattern:     "/user/",
			HandlerFunc: UpdateUser,
			Protected:   true,
		},
		model.Route{
			Name:        "DeleteUserByID",
			Method:      "DELETE",
			Pattern:     "/user/{id}",
			HandlerFunc: DeleteUserByID,
			Protected:   true,
		},
		model.Route{
			Name:        "GetUserLogin",
			Method:      "POST",
			Pattern:     "/login",
			HandlerFunc: GetUserLogin,
			Protected:   false,
		},
		model.Route{
			Name:        "GetUserLogout",
			Method:      "POST",
			Pattern:     "/logout",
			HandlerFunc: GetUserLogout,
			Protected:   true,
		},
	},
}

// GetUserByID by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user model.MUser
	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = repository.GetUserByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (model.MUser{}) == user {
		util.ResponseWithJSON(w, http.StatusNotFound, user)
	} else {
		util.ResponseWithJSON(w, http.StatusOK, user)
	}
}

// GetUserLogin for login purpose
func GetUserLogin(w http.ResponseWriter, r *http.Request) {

	var creds CredentialsLogin
	var user model.MUser
	var err error

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = repository.GetUserLogin(creds.Username, creds.Password)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	jwt, err := util.CreateToken(user)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = util.SaveToRedis(user.ID, jwt)
	if err != nil {
		ad := &util.AccessDetails{
			AccessUUID: jwt.AccessUUID,
			UserID:     user.ID,
		}
		util.DeleteToken(ad)
		util.ResponseWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, jwt)
}

// GetUserLogout ...
func GetUserLogout(w http.ResponseWriter, r *http.Request) {

	accessDetails, err := util.ExtractFromRedis(r)
	if err != nil {
		util.ResponseWithError(w, http.StatusUnauthorized, "Verify Token failure. Reason: "+err.Error())
		return
	}

	err = util.DeleteToken(accessDetails)
	if err != nil {
		util.ResponseWithError(w, http.StatusUnauthorized, "Delete Token failure. Reason: "+err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, map[string]string{"message": "Successfully logged out!"})
}

// GetUsers data
func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []model.MUser
	users, err := repository.GetUserAll()
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, users)
}

// CreateUser post
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.MUser

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &user)

	user, err = repository.CreateUser(user)
	if err != nil {
		util.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusCreated, user)
}

// UpdateUser update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user model.MUser

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &user)

	usr, err := repository.UpdateUser(user)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, usr)
}

// DeleteUserByID ...
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	var user model.MUser

	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	err = repository.DeleteUserByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusNoContent, user)
}
