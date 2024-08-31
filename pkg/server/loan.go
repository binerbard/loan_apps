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


func LoanServer(app *fiber.App) {
	// TODO: implement
	db := database.NewDatabaseConnection()
	loanRepo := repositories.NewLoanRepository(db)
	tenorRepo := repositories.NewTenorRepository(db)
	loanServer := services.NewtLoanService(loanRepo,tenorRepo)
	loanUseCase := usecase.NewULoanUsecase(loanRepo,loanServer)
	loanHandler := handler.NewLoanHandler(loanUseCase)

	routes.RegisterLoanRoutes(app, loanHandler)
}