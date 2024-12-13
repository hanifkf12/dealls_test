package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/pkg/config"
	"github.com/hanifkf12/hanif_skeleton/pkg/jwtx"
	"os"
)

func JWTMiddleware(xCtx *fiber.Ctx, conf *config.Config) appctx.Response {

	ctx := xCtx.Context()

	bearerToken := xCtx.Get(fiber.HeaderAuthorization)

	if bearerToken == "" {
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithErrors("Unauthorized")
	}
	token := bearerToken[7:]

	fmt.Println(token)

	secretKey, err := os.ReadFile("./secret/secret.key")
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	claims, err := jwtx.ValidateJWT(token, secretKey)
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithErrors(err.Error())
	}

	err = claims.Valid()
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithErrors(err.Error())
	}
	ctx.SetUserValue("user_id", claims.UserID)
	ctx.SetUserValue("email", claims.Email)

	xCtx.SetUserContext(ctx)
	return *appctx.NewResponse().WithCode(fiber.StatusOK)
}
