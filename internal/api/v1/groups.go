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
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(addedGroup)
	if err != nil {
		g.lg.Error("AddGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (g *Groups) EditGroup(writer http.ResponseWriter, request *http.Request) {
	var data GroupDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		g.lg.Error("EditGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if (data.ID == 0) || (data.Name == "") {
		g.lg.Error("EditGroup: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var group = model.Group{
		ID:   data.ID,
		Name: data.Name,
	}
	writer.Header().Set("Content-Type", "application/json")
	err = g.groupsRepo.EditGroup(request.Context(), group)
	if err != nil {
		g.lg.Error("EditGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func (g *Groups) ListAllGroups(writer http.ResponseWriter, request *http.Request) {
	groups, err := g.groupsRepo.ListAllGroups(request.Context())
	if err != nil {
		g.lg.Error("ListAllGroups", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(groups)
	if err != nil {
		g.lg.Error("ListAllGroups", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func (g *Groups) DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	var data GroupDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		g.lg.Error("DeleteGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if data.ID == 0 {
		g.lg.Error("DeleteGroup: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = g.groupsRepo.DeleteGroupbyID(request.Context(), data.ID)
	if err != nil {
		g.lg.Error("DeleteGroup", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (g *Groups) GetGroupID(writer http.ResponseWriter, request *http.Request) {
	var data GroupDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		g.lg.Error("GetGroupID", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if data.Name == "" {
		g.lg.Error("GetGroupID: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id, err := g.groupsRepo.GetGroupIDbyName(request.Context(), data.Name)
	if err != nil {
		g.lg.Error("GetGroupID", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var group = GroupDTO{ID: id}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(group)
	if err != nil {
		g.lg.Error("GetGroupID", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (g *Groups) GetGroupSummary(writer http.ResponseWriter, request *http.Request) {
	var data GroupDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		g.lg.Error("GetGroupSummary", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if data.ID == 0 {
		g.lg.Error("GetGroupSummary: field is empty")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	members, err := g.groupsRepo.GetGroupMembersbyID(request.Context(), data.ID)
	if err != nil {
		g.lg.Error("GetGroupSummary", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(members)
	if err != nil {
		g.lg.Error("GetGroupSummary", zap.Error(err))
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
