package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all Loan routes
var RouterLoanList map[string]string = map[string]string{
	"main":"/loan",
	"store":"/store",
	"update":"/update/:id",
	"delete":"/delete/:id",
	"find":"/find/:id",
	"list":"/list",
	"customer":"/customer/:id",
}


func RegisterLoanRoutes(r fiber.Router, handler *handler.LoanHandler) {
	router := r.Group(RouterLoanList["main"])
	router.Post(RouterLoanList["store"], handler.Store)
	router.Put(RouterLoanList["update"], handler.Update)
	router.Delete(RouterLoanList["delete"], handler.Delete)
	router.Get(RouterLoanList["find"], handler.Find)
	router.Get(RouterLoanList["list"], handler.List)
	router.Get(RouterLoanList["customer"], handler.FindByCustomerID)
}
