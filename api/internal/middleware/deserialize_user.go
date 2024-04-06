package middleware

import (
	"Jwtwithecdsa/api/internal/utils"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var access_token string
		authorization := ctx.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer ") {
			access_token = strings.TrimPrefix(authorization, "Bearer ")
		}
		if access_token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
		}
		tokenClaims, err := utils.VerifyToken(access_token, os.Getenv("PUBLIC_ACCESS_KEY"))
		if err != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		claims, ok := tokenClaims.Claims.(jwt.MapClaims)
		if ok && tokenClaims.Valid {
			if int64(claims["exp"].(float64)) < time.Now().Unix() {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "your session has expired"})
			}
		}
		userid := claims["userid"].(string)
		if userid == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "user not found"})
		}
		ctx.Locals("userid", userid)
		return ctx.Next()
	}
}
