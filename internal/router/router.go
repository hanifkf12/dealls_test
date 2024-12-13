package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanifkf12/hanif_skeleton/internal/appctx"
	"github.com/hanifkf12/hanif_skeleton/internal/bootstrap"
	"github.com/hanifkf12/hanif_skeleton/internal/handler"
	"github.com/hanifkf12/hanif_skeleton/internal/middleware"
	"github.com/hanifkf12/hanif_skeleton/internal/repository/home"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase"
	"github.com/hanifkf12/hanif_skeleton/internal/usecase/contract"
	"github.com/hanifkf12/hanif_skeleton/pkg/config"
	"github.com/hanifkf12/hanif_skeleton/pkg/logger"
)

type router struct {
	cfg   *config.Config
	fiber fiber.Router
}

func (rtr *router) handle(hfn httpHandlerFunc, svc contract.UseCase, mdws ...middleware.MiddlewareFunc) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if rm := middleware.FilterFunc(rtr.cfg, ctx, mdws); rm.Code != fiber.StatusOK {
			// return response base on middleware
			res := *appctx.NewResponse().
				WithCode(rm.Code).
				WithErrors(rm.Errors).
				WithMessage(rm.Message)
			return rtr.response(ctx, res)
		}

		resp := hfn(ctx, svc, rtr.cfg)
		return rtr.response(ctx, resp)
	}
}

func (rtr *router) response(ctx *fiber.Ctx, resp appctx.Response) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(resp.Code).Send(resp.Byte())
}

type test struct {
}

func (t *test) Serve(data appctx.Data) appctx.Response {
	logger.Info("test")
	ctctx := data.FiberCtx.UserContext()
	logger.Info(ctctx.Value("user_id"))
	logger.Info(ctctx.Value("email"))
	return *appctx.NewResponse().WithMessage("test").WithCode(fiber.StatusOK)
}

func (rtr *router) Route() {
	db := bootstrap.RegistryDatabase(rtr.cfg)
	homeRepo := home.NewHomeRepository(db)

	healthUseCase := usecase.NewHealth(homeRepo)
	testSvc := &test{}
	rtr.fiber.Get("/health", rtr.handle(
		handler.HttpRequest,
		healthUseCase,
	))

	rtr.fiber.Get("/ping", rtr.handle(
		handler.HttpRequest,
		testSvc,
		middleware.JWTMiddleware,
	))

}

func NewRouter(cfg *config.Config, fiber fiber.Router) Router {
	return &router{
		cfg:   cfg,
		fiber: fiber,
	}
}
