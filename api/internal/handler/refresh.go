package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) RefreshAccessToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		refreshtoken := strings.Split(authHeader, "Bearer ")
		if len(refreshtoken) != 2 {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "token is invalid or session has expired",
			})
		}
		response, err := h.ctrl.RefreshAccessToken(refreshtoken[1])
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"success":      true,
			"Accesstoken":  response.AccessToken,
			"Refreshtoken": response.RefreshToken,
		})
	}
}
