package rest

import (
	"loan_apps/pkg/server"
	"loan_apps/config/setting"
	"github.com/gofiber/fiber/v2"
)


func Main(){
	app:= fiber.New()

	server.CustomerServer(app)
	server.TenorServer(app)
	server.LoanServer(app)
	server.TransactionServer(app)
	server.BalanceServer(app)
	app.Listen(":"+setting.GetConfig().AppPort)
}
