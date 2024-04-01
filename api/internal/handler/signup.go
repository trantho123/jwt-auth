package handler

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var signUpInput model.SignUpInput
		if err := ctx.BodyParser(&signUpInput); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		if signUpInput.UserName == "" || signUpInput.Password == "" || signUpInput.Email == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing input",
			})
		}
		if err := utils.IsValidEmail(signUpInput.Email); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		err := h.ctrl.SighUp(&signUpInput)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"Message": "Please verify your account by clicking on the link in the email we've sent you",
		})
	}
}
