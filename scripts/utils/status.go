package utils

import "github.com/gofiber/fiber/v2"

// BadRequest returns a 400 Bad Request response with the given message.
func BadRequest(c *fiber.Ctx, msg fiber.Map) error {
	msg["code"] = 400
	msg["status"] = "Bad Request"
	return c.Status(fiber.StatusBadRequest).JSON(msg)
}

// NotFound returns a 404 Not Found response with the given message.
func NotFound(c *fiber.Ctx, msg fiber.Map) error {
	msg["code"] = 404
	msg["status"] = "Not Found"
	return c.Status(fiber.StatusNotFound).JSON(msg)
}

// InternalServerError returns a 500 Internal Server Error response with the given message.
func InternalServerError(c *fiber.Ctx, msg fiber.Map) error {
	msg["code"] = 500
	msg["status"] = "Internal Server Error"
	return c.Status(fiber.StatusInternalServerError).JSON(msg)
}

// Success returns a 200 Success response with the given message.
func Success(c *fiber.Ctx, msg fiber.Map) error {
	msg["code"] = 200
	msg["status"] = "Success"
	return c.Status(fiber.StatusOK).JSON(msg)
}

// UnAuthorized returns a 401 Unauthorized response with the given message.
func UnAuthorized(c *fiber.Ctx, msg fiber.Map) error {
	msg["code"] = 401
	msg["status"] = "UnAuthorized"
	return c.Status(fiber.StatusUnauthorized).JSON(msg)
}
