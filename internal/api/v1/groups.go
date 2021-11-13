package v1

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"olaf/internal/model"
	"olaf/internal/repository"
)

type Groups struct {
	groupsRepo repository.Groups
	usersRepo  repository.Users
	lg         *zap.Logger
}

func NewGroups(groupsRepo repository.Groups, usersRepo repository.Users, lg *zap.Logger) *Groups {
	return &Groups{groupsRepo: groupsRepo, usersRepo: usersRepo, lg: lg}
}

func (g *Groups) AddGroup(writer http.ResponseWriter, request *http.Request) {
	var data GroupDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		g.lg.Error("AddGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if data.Name == "" {
		g.lg.Error("AddGroup: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var group = model.Group{
		Name:     data.Name,
		IsActive: data.IsActive,
	}
	var addedGroup GroupDTO
	addedGroup.ID, err = g.groupsRepo.AddGroup(request.Context(), group)
	if err != nil {
		g.lg.Error("AddGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(addedGroup)
	if err != nil {
		g.lg.Error("AddGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
}

func (g *Groups) EditGroup(writer http.ResponseWriter, request *http.Request) {

}
func (g *Groups) ListAllGroups(writer http.ResponseWriter, request *http.Request) {

}
func (g *Groups) DeleteGroup(writer http.ResponseWriter, request *http.Request) {

}
