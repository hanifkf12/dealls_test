package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/presentation"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type signup struct {
	userRepo    repository.User
	profileRepo repository.Profile
}

func (s *signup) Serve(data appctx.Data) appctx.Response {
	var (
		lf      = logger.NewFields(logger.EventName("v1.SignUp"))
		ctx     = data.FiberCtx.UserContext()
		payload presentation.Signup
	)

	err := data.FiberCtx.BodyParser(&payload)
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnprocessableEntity).WithErrors(err.Error())
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	payload.Password = string(bytes)
	userId, err := s.userRepo.Create(ctx, entity.User{
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	})
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	_, err = s.profileRepo.Create(ctx, entity.Profile{
		UserID:    int(userId),
		Name:      payload.Name,
		Avatar:    payload.Avatar,
		Age:       payload.Age,
		Gender:    payload.Gender,
		Bio:       payload.Bio,
		Location:  "",
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	})
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("OK")
}

func NewSignUp(userRepo repository.User, profileRepo repository.Profile) contract.UseCase {
	return &signup{
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}
