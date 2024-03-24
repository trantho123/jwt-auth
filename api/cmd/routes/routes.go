package routes

import (
	"Jwtwithecdsa/api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	h handler.Handler
}

func (r Route) Routes(rtr *fiber.App) {
	rtr.Post("/test/signup", r.h.SignUp())
}
