package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all customer routes
var RouterCustomerList map[string]string = map[string]string{
	"main":"/customer",
	"register":"/register",
}


func RegisterCustomerRoutes(r fiber.Router, handler *handler.CustomerHandler) {
	router := r.Group(RouterCustomerList["main"])
	router.Post(RouterCustomerList["register"], handler.Registration)
}
