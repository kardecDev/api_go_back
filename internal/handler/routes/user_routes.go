package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/kardecDev/api_go_back/config/env"
	"github.com/kardecDev/api_go_back/internal/handler/midleware"
	"github.com/kardecDev/api_go_back/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Use(midleware.LoggerData)

	router.Post("/user", h.CreateUser)

	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		//rotas protegidas para usuários
		r.Patch("/user", h.UpdateUser)
		r.Get("/user", h.GetUserByID)
		r.Delete("/user", h.DeleteUser)
		r.Get("/user/list-all", h.FindManyUsers)
		r.Patch("/user/password", h.UpdateUserPassword)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
