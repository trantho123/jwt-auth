package routes

import (
	"Jwtwithecdsa/api/internal/handler"
)

func New(h handler.Handler) Route {
	return Route{h: h}
}
