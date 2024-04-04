package handler

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) LoginVerify() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var verifyOTP model.VerifyOTP
		if err := ctx.BodyParser(&verifyOTP); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		if verifyOTP.Email == "" || verifyOTP.OTPCode == "" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Missing Username or Password",
			})
		}
		if err := utils.IsValidEmail(verifyOTP.Email); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		response, err := h.ctrl.LoginVerify(&verifyOTP)
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
