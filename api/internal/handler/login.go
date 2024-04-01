package handler

import (
	"Jwtwithecdsa/api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var loginInput model.LoginInput
		if err := ctx.BodyParser(&loginInput); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		if loginInput.UserName == "" || loginInput.Password == "" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Missing Username or Password",
			})
		}
		err := h.ctrl.Login(&loginInput)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"Message": "Please enter the OTP code that we have sent to your email to proceed",
		})
	}
}
