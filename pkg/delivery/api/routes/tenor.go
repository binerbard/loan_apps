package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all tenor routes
var RouterTenorList map[string]string = map[string]string{
	"main":"/tenor",
	"store":"/store",
	"update":"/update/:id",
	"delete":"/delete/:id",
	"find":"/find/:id",
	"list":"/list",
	"customer":"/customer/:id",
}


func RegisterTenorRoutes(r fiber.Router, handler *handler.TenorHandler) {
	router := r.Group(RouterTenorList["main"])
	router.Post(RouterTenorList["store"], handler.Store)
	router.Get(RouterTenorList["list"], handler.List)
	router.Get(RouterTenorList["find"], handler.Find)
	router.Put(RouterTenorList["update"], handler.Update)
	router.Delete(RouterTenorList["delete"], handler.Delete)
	router.Get(RouterTenorList["customer"], handler.FindByCustomerID)
}
