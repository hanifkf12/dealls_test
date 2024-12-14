package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/presentation"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/jwtx"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	userRepo repository.User
}

func (l *login) Serve(data appctx.Data) appctx.Response {
	var (
		lf      = logger.NewFields(logger.EventName("v1.Login"))
		ctx     = data.FiberCtx.UserContext()
		payload presentation.Signup
	)

	err := data.FiberCtx.BodyParser(&payload)
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnprocessableEntity).WithErrors(err.Error())
	}

	user, err := l.userRepo.FindByEmail(ctx, payload.Email)
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithErrors(err.Error())
	}

	jwt, err := jwtx.GenerateJWT(user.Email, user.ID)
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())

	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("OK").WithData(presentation.LoginResponse{Token: jwt})
}

func NewLogin(userRepo repository.User) contract.UseCase {
	return &login{
		userRepo: userRepo,
	}
}
