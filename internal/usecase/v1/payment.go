package v1

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/presentation"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
	"time"
)

type payment struct {
	transactionRepo repository.Transaction
}

func (p *payment) Serve(data appctx.Data) appctx.Response {
	var (
		lf      = logger.NewFields(logger.EventName("v1.Payment"))
		ctx     = data.FiberCtx.UserContext()
		payload presentation.Payment
	)

	err := data.FiberCtx.BodyParser(&payload)
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusUnprocessableEntity).WithErrors(err.Error())
	}

	duration, err := p.getPaymentDuration(payload.PackageName)
	if err != nil {
		return appctx.Response{}
	}
	_, err = p.transactionRepo.CreateTransaction(ctx, entity.Transaction{
		UserID:      payload.UserID,
		Price:       payload.Amount,
		PackageType: payload.PackageName,
		ValidUntil:  time.Now().Add(duration),
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	})
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}
	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("OK")
}

func (p *payment) getPaymentDuration(duration string) (time.Duration, error) {
	switch duration {
	case "7_day":
		return 7 * 24 * time.Hour, nil
	case "1_week":
		return 7 * 24 * time.Hour, nil
	case "1_month":
		return 30 * 24 * time.Hour, nil

	default:
		return 0, errors.New("invalid duration")

	}
}

func NewPayment(transactionRepo repository.Transaction) contract.UseCase {
	return &payment{
		transactionRepo: transactionRepo,
	}
}
