package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all customer routes
var RouterCustomerList map[string]string = map[string]string{
	"main":"/customer",
	"store":"/store",
	"update":"/update/:id",
	"delete":"/delete/:id",
	"find":"/find/:id",
	"list":"/list",
	"show":"/show/:filename",
}


func RegisterCustomerRoutes(r fiber.Router, handler *handler.CustomerHandler) {
	router := r.Group(RouterCustomerList["main"])
	router.Post(RouterCustomerList["store"], handler.Store)
	router.Get(RouterCustomerList["show"],handler.ShowImage)
	router.Put(RouterCustomerList["update"], handler.Update)
	router.Delete(RouterCustomerList["delete"], handler.Delete)
	router.Get(RouterCustomerList["list"], handler.List)
	router.Get(RouterCustomerList["find"], handler.Find)
}
