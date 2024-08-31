package rest

import (
	"loan_apps/pkg/server"

	"github.com/gofiber/fiber/v2"
)


func Main(){
	app:= fiber.New()

	server.CustomerServer(app)
	server.TenorServer(app)
	server.LoanServer(app)
	server.TransactionServer(app)
	app.Listen(":3000")
}