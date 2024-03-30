package routes

import (
	"Jwtwithecdsa/api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	h handler.Handler
}

func (r Route) Routes(rtr *fiber.App) {
	rtr.Post("/signup", r.h.SignUp())
	rtr.Post("/signup/verifyemail", r.h.SignUpVerifyEmail())
}
