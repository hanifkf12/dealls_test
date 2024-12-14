package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
)

type getProfiles struct {
	profileRepo repository.Profile
}

func (g *getProfiles) Serve(data appctx.Data) appctx.Response {
	var (
		lf = logger.NewFields(
			logger.EventName("v1.GetProfiles"),
		)
		ctx    = data.FiberCtx.UserContext()
		userID = ctx.Value("user_id").(int)
		email  = ctx.Value("email").(string)
	)

	lf.Append(logger.Any("userID", userID))
	lf.Append(logger.Any("email", email))

	currentProfile, err := g.profileRepo.FindByUsersID(ctx, userID)
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	if currentProfile == nil {
		logger.Error("Profile not found", lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusNotFound).WithErrors("Profile not found")
	}

	profiles, err := g.profileRepo.FindAll(ctx, userID, currentProfile.Gender)
	if err != nil {
		logger.Error(err.Error(), lf...)
		return *appctx.NewResponse().WithCode(fiber.StatusInternalServerError).WithErrors(err.Error())
	}

	logger.Info("OK", lf...)

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithData(profiles).WithMessage("Success")
}

func NewGetProfiles(profileRepo repository.Profile) contract.UseCase {
	return &getProfiles{
		profileRepo: profileRepo,
	}
}
