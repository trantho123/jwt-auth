package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetMe() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userid := ctx.Locals("userid").(string)
		userRes, err := h.ctrl.GetMe(userid)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"User":    userRes,
		})
	}
}
