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
	var data UserDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if (data.Name == "") || (data.Password == "") || (data.Email == "") || (data.Group == 0) {
		u.lg.Error("Adduser: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var user = model.User{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		IsActive: data.IsActive,
		Group:    data.Group,
	}

	var addedUser UserDTO
	addedUser.ID, err = u.usersRepo.AddUser(request.Context(), user)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(addedUser)
	if err != nil {
		u.lg.Error("AddUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
}

func (u *Users) EditUser(writer http.ResponseWriter, request *http.Request) {
	var data UserDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		u.lg.Error("EditUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if data.ID == 0 {
		u.lg.Error("EditUser: field id is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var user = model.User{
		ID:       data.ID,
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		IsActive: data.IsActive,
		Group:    data.Group,
	}
	err = u.usersRepo.EditUser(request.Context(), user)
	if err != nil {
		u.lg.Error("EditUser", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
}

func (u *Users) ListAllUsers(writer http.ResponseWriter, request *http.Request) {

}

func (u *Users) DeleteUser(writer http.ResponseWriter, request *http.Request) {

}
