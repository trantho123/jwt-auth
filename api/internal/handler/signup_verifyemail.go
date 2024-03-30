package handler

import (
	"Jwtwithecdsa/api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignUpVerifyEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var verify model.VerifyEmail
		if err := ctx.BodyParser(&verify); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		err := h.ctrl.SignUpVerifyEmail(&verify)
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
