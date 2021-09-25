package main

import (
	"net/http"

	"github.com/gopackager/micro-based/config"
	"github.com/gopackager/micro-based/helper"
	"github.com/gopackager/micro-based/helper/asynq/client"
	"github.com/gopackager/micro-based/repositories"
	"github.com/gopackager/micro-based/transports/api/rest"
	"github.com/gopackager/micro-based/usecases"
)

func main() {
	app := config.Application{
		Mode:    helper.Env("APP_ENV", "local"),
		Version: helper.Env("APP_VERSION", "v1"),
		Port:    helper.Env("APP_PORT", "8088"),
	}
	db := config.Database{
		Username: helper.Env("MYSQL_USERNAME", "root"),
		Password: helper.Env("MYSQL_PASSWORD", "password"),
		Host:     helper.Env("MYSQL_HOST", "localhost"),
		Port:     helper.Env("MYSQL_PORT", "3306"),
		DBName:   helper.Env("MYSQL_NAME", "job"),
		DBDriver: helper.Env("MYSQL_DRIVER", "mysql"),
	}

	queue := client.New(helper.Env("REDIS_HOST"), helper.Env("REDIS_PORT"), helper.Env("REDIS_PASSWORD"))
	defer queue.Close()

	router := config.New(app)
	mysql := config.MySQL(db)

	repo := repositories.New(mysql, queue)
	usecase := usecases.New(repo)
	transport := rest.New(usecase)
	router.RegisterRoutes(route(transport)).Starts()
}

func route(transport rest.Transports) []config.Routing {
	return []config.Routing{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: transport.Create,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: transport.Read,
		},
		{
			Method:  http.MethodGet,
			Path:    "/user/:id",
			Handler: transport.Detail,
		},
		{
			Method:  http.MethodPut,
			Path:    "/user/:id",
			Handler: transport.Update,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/user/:id",
			Handler: transport.Delete,
		},
	}
}
