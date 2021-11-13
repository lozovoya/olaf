package v1

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"olaf/internal/model"
	"olaf/internal/repository"
)

type Users struct {
	usersRepo repository.Users
	lg        *zap.Logger
}

func NewUsers(usersRepo repository.Users, lg *zap.Logger) *Users {
	return &Users{usersRepo: usersRepo, lg: lg}
}

func (u *Users) AddUser(writer http.ResponseWriter, request *http.Request) {
	var data model.User
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user, err := u.usersRepo.AddUser(request.Context(), data)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
}

func (u *Users) EditUser(writer http.ResponseWriter, request *http.Request) {

}

func (u *Users) ListAllUsers(writer http.ResponseWriter, request *http.Request) {

}

func (u *Users) DeleteUser(writer http.ResponseWriter, request *http.Request) {

}
