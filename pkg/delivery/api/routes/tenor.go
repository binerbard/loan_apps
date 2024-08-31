package routes

import (
	"loan_apps/pkg/delivery/api/handler"

	"github.com/gofiber/fiber/v2"
)

// RouterAuthList list of all tenor routes
var RouterTenorList map[string]string = map[string]string{
	"main":"/tenor",
	"store":"/store",
}


func RegisterTenorRoutes(r fiber.Router, handler *handler.TenorHandler) {
	router := r.Group(RouterTenorList["main"])
	router.Post(RouterTenorList["store"], handler.Store)
}
