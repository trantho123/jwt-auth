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
	rtr.Get("/signup/verify/:id/:code", r.h.SignUpVerifyEmail())
	rtr.Post("/login", r.h.Login())
	rtr.Post("/login/verify", r.h.LoginVerify())
}
