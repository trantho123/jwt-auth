package handler

import (
	"Jwtwithecdsa/api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var SignUpInput model.SignUpInput
		if err := ctx.BodyParser(&SignUpInput); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		err := h.ctrl.SighUp(&SignUpInput)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
		})
	}
}
