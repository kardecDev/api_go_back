package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	logger "github.com/kardecDev/api_go_back/config"
	"github.com/kardecDev/api_go_back/config/env"
	_ "github.com/kardecDev/api_go_back/docs"
	"github.com/kardecDev/api_go_back/internal/database"
	"github.com/kardecDev/api_go_back/internal/database/sqlc"
	"github.com/kardecDev/api_go_back/internal/handler/routes"
	"github.com/kardecDev/api_go_back/internal/handler/userhandler"
	"github.com/kardecDev/api_go_back/internal/repository/userrepository"
	"github.com/kardecDev/api_go_back/internal/service/userservice"
)

func main() {
	logger.InitLogger()

	slog.Info("Starting API") //Level logs: Info, Error, Debug, Warm
	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load enviroment variables", slog.String("error", err.Error()), slog.String("pakage", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("Error to connect to database",
			"err", err,
			slog.String("package", "main"),
		)
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	//user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	//init routes
	routes.InitUserRoutes(router, newUserHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("Server running on port %s", port))
	err = http.ListenAndServe(port, router)

	if err != nil {
		slog.Error("Error to start server", slog.String("error", err.Error()), slog.String("package", "main"))
	}
}
