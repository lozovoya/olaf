package v1

import (
	"go.uber.org/zap"
	"net/http"
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

}
func (g *Groups) EditGroup(writer http.ResponseWriter, request *http.Request) {

}
func (g *Groups) ListAllGroups(writer http.ResponseWriter, request *http.Request) {

}
func (g *Groups) DeleteGroup(writer http.ResponseWriter, request *http.Request) {

}
