package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all tenor routes
var RouterTransactionList map[string]string = map[string]string{
	"main":"/transaction",
	"store":"/store",
}


func RegisterTransactionRoutes(r fiber.Router, handler *handler.TransactionHandler) {
	router := r.Group(RouterTransactionList["main"])
	router.Post(RouterTransactionList["store"], handler.Store)
}
