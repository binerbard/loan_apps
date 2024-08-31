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


func TransactionServer(app *fiber.App) {
	// TODO: implement
	db := database.NewDatabaseConnection()
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionServer := services.NewtTransactionService(transactionRepo)
	transactionUseCase := usecase.NewUTransactionUsecase(transactionRepo,transactionServer)
	transactionHandler := handler.NewTransactionHandler(transactionUseCase)

	routes.RegisterTransactionRoutes(app, transactionHandler)
}