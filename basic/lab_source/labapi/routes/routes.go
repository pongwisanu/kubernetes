package routes

import (
	"labapi/handlers"
	"labapi/repositories"
	"labapi/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Register(api *fiber.App) {
	dsn := "postgres://dbadmin:dbadmin@labdb-service:5432/postgres"
	connConfig, _ := pgx.ParseConfig(dsn)
	pgxdb := stdlib.OpenDB(*connConfig)
	db := sqlx.NewDb(pgxdb, "pgx")

	repo := repositories.NewUserRepositoryDb(db)
	service := services.NewUserService(repo)
	u_handler := handlers.NewUserHandler(service)

	api.Get("/", handlers.GetIndex)

	api.Get("/users", u_handler.GetUsers)
	api.Get("/user/:id", u_handler.GetUser)
	api.Post("/user/add", u_handler.AddUser)
}
