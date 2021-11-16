package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	v1 "olaf/internal/api/v1"
)

func NewRouter(mux *chi.Mux, usersController *v1.Users, groupsController *v1.Groups, lg *zap.Logger) chi.Mux {
	mux.Use(middleware.Logger)
	mux.Route("/api/v1", func(router chi.Router) {
		RouterUsers(router, usersController)
		RouterGroups(router, groupsController)
	})
	lg.Info("Router is activated")
	return *mux
}

func RouterUsers(router chi.Router, usersController *v1.Users) {
	router.Post("/users", usersController.AddUser)
	router.Post("/users/id", usersController.GetUsersID)
	router.Put("/users", usersController.EditUser)
	router.Put("/users/delete", usersController.DeleteUser)
	router.Get("/users", usersController.ListAllUsers)
}

func RouterGroups(router chi.Router, groupsController *v1.Groups) {
	router.Post("/groups", groupsController.AddGroup)
	router.Put("/groups", groupsController.EditGroup)
	router.Put("/groups/delete", groupsController.DeleteGroup)
	router.Get("/groups", groupsController.ListAllGroups)
	router.Post("/groups/id", groupsController.GetGroupID)
	router.Post("/groups/sum", groupsController.GetGroupSummary)
}
