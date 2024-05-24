package main

import (
	"log/slog"

	logger "github.com/kardecDev/api_go_back/config"
)

// create an struct for management user
type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

// create a function for handle logs of user
func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "Hidden"),
	)
}

func main() {
	logger.InitLogger()

	//Instance of user
	user := user{
		Name:     "Allan Kardec",
		Age:      47,
		Password: "teste123",
	}
	slog.Info("Starting API") //Level logs: Info, Error, Debug, Warm
	slog.Info("Creating user", "user", user.LogUser())

}
