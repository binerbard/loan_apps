package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all Loan routes
var RouterLoanList map[string]string = map[string]string{
	"main":"/loan",
	"store":"/store",
}


func RegisterLoanRoutes(r fiber.Router, handler *handler.LoanHandler) {
	router := r.Group(RouterLoanList["main"])
	router.Post(RouterLoanList["store"], handler.Store)
}
