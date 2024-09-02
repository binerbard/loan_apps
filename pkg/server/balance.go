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


func BalanceServer(app *fiber.App) {
	// TODO: implement
	db := database.NewDatabaseConnection()
	balanceRepo := repositories.NewBalanceRepository(db)
	balanceServer := services.NewBalanceService(balanceRepo)
	balanceUseCase := usecase.NewBalanceUsecase(balanceRepo,balanceServer)
	balanceHandler := handler.NewBalanceHandler(balanceUseCase)

	routes.RegisterBalanceRoutes(app, balanceHandler)
}