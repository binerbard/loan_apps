package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all Balance routes
var RouterBalanceList map[string]string = map[string]string{
	"main":"/balance",
	"find":"/find/:id",
	"list":"/list",
	"customer":"/customer/:id",
}


func RegisterBalanceRoutes(r fiber.Router, handler *handler.BalanceHandler) {
	router := r.Group(RouterBalanceList["main"])
	router.Get(RouterBalanceList["list"], handler.List)
	router.Get(RouterBalanceList["find"], handler.Find)
	router.Get(RouterBalanceList["customer"], handler.FindByCustomerID)
}
