package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/presentation"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
	"time"
)

type swipe struct {
	profileRepo     repository.Profile
	swipeRepo       repository.Swipe
	transactionRepo repository.Transaction
}

func (s *swipe) Serve(data appctx.Data) appctx.Response {
	var (
		lf = logger.NewFields(
			logger.EventName("v1.Swipe"),
		)
		ctx     = data.FiberCtx.UserContext()
		userID  = ctx.Value("user_id").(int)
		email   = ctx.Value("email").(string)
		payload presentation.Swipe
	)
	err := data.FiberCtx.BodyParser(&payload)
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnprocessableEntity).WithErrors(err.Error())
	}
	lf.Append(logger.Any("userID", userID))
	lf.Append(logger.Any("email", email))
	lf.Append(logger.Any("swipeType", payload.SwipeType))

	isPremium, err := s.transactionRepo.CheckPremiumStatus(ctx, userID)
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	if !isPremium {
		limit, err := s.swipeRepo.IsLimit(ctx, userID)
		if err != nil {
			logger.Error(err.Error(), lf...)
			return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
		}
		if limit {
			logger.Error("You have reached the limit of swiping", lf...)
			return *appctx.NewResponse().WithCode(fiber.StatusTooManyRequests).WithErrors("You have reached the limit of swiping")
		}
	}

	switch payload.SwipeType {
	case "like":
		_, err := s.swipeRepo.SwipeRight(ctx, entity.Swipe{
			UserID:    userID,
			ProfileID: payload.ProfileID,
			SwipeType: 1,
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		})
		if err != nil {
			logger.Error(err.Error(), lf...)
			return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
		}
	case "dislike":
		_, err := s.swipeRepo.SwipeRight(ctx, entity.Swipe{
			UserID:    userID,
			ProfileID: payload.ProfileID,
			SwipeType: 0,
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		})
		if err != nil {
			logger.Error(err.Error(), lf...)
			return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
		}
	}

	logger.Info("Swiped", lf...)
	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("OK").WithMessage("Swiped")

}

func NewSwipe(profileRepo repository.Profile, swipeRepo repository.Swipe, transaction repository.Transaction) contract.UseCase {
	return &swipe{
		profileRepo:     profileRepo,
		swipeRepo:       swipeRepo,
		transactionRepo: transaction,
	}
}
