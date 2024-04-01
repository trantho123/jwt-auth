package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignUpVerifyEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		code := ctx.Params("code")
		if id == "" || code == "" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Missing ID or code parameter",
			})
		}
		err := h.ctrl.SignUpVerifyEmail(id, code)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"Message": "Email verified successfully",
		})
	}
}
