package handler

import (
	"Jwtwithecdsa/api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user model.User
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		err := h.ctrl.SighUp(&user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not sign up user",
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
		})
	}
}
