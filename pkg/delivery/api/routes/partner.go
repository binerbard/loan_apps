package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

var RouterPartnerList map[string]string = map[string]string{
	"main":"/partner",
	"store":"/store",
	"update":"/update/:id",
	"delete":"/delete/:id",
	"find":"/find/:id",
	"list":"/list",
}


func RegisterPartnerRoutes(r fiber.Router, handler *handler.PartnerHandler) {
	router := r.Group(RouterPartnerList["main"])
	router.Post(RouterPartnerList["store"], handler.Store)
	router.Put(RouterPartnerList["update"], handler.Update)
	router.Delete(RouterPartnerList["delete"], handler.Delete)
	router.Get(RouterPartnerList["find"], handler.Find)
	router.Get(RouterPartnerList["list"], handler.List)
}