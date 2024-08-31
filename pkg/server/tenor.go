package server

import (
	"loan_apps/datasource/domain/usecase"
	"loan_apps/datasource/repositories"
	"loan_apps/internal/infrastructure/database"
	"loan_apps/internal/services"
	"loan_apps/pkg/delivery/api/handler"
	"loan_apps/pkg/delivery/api/routes"

	"github.com/gofiber/fiber/v2"
)


func TenorServer(app *fiber.App) {
	// TODO: implement
	db := database.NewDatabaseConnection()
	tenorRepo := repositories.NewTenorRepository(db)
	tenorServer := services.NewtTenorService(tenorRepo)
	tenorUseCase := usecase.NewUTenorUsecase(tenorRepo,tenorServer)
	tenorHandler := handler.NewTenorHandler(tenorUseCase)

	routes.RegisterTenorRoutes(app, tenorHandler)
}