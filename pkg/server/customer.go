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


func CustomerServer(app *fiber.App) {
	// TODO: implement
	db := database.NewDatabaseConnection()
	customerRepo := repositories.NewCustomerRepository(db)
	customerServer := services.NewCustomerService(customerRepo)
	customerUseCase := usecase.NewUCustomerUsecase(customerRepo,customerServer)
	customerHandler := handler.NewCustomerHandler(customerUseCase)

	routes.RegisterCustomerRoutes(app, customerHandler)
}