package authHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
	"github.com/hifat/con-q-api/internal/pkg/token"
)

type AuthHandler struct {
	cfg config.AppConfig

	authService authDomain.IAuthService
}

func New(cfg config.AppConfig, authService authDomain.IAuthService) AuthHandler {
	return AuthHandler{cfg, authService}
}

// @Summary		Register
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ReqRegister
// @Success		409 {object} errorDomain.Response "Duplicate record"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/register [post]
// @Param		Body body authDomain.ReqRegister true "Register request"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req authDomain.ReqRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.ValidateFormErr(ctx, err)
		return
	}

	res, err := h.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary		Login
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResToken
// @Success		401 {object} errorDomain.Response "Invalid Credentials"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/login [post]
// @Param		Body body authDomain.ReqLogin true "Login request"
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req authDomain.ReqLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.ValidateFormErr(ctx, err)
		return
	}

	req.Agent = ctx.Request.UserAgent()
	req.ClientIP = ctx.ClientIP()

	res, err := h.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary		Refresh Token
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResToken
// @Success		401 {object} errorDomain.Response "Invalid Credentials"
// @Success		401 {object} errorDomain.Response "Revoked Token"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/refresh-token [post]
// @Param		Body body authDomain.ReqLogin true "Login request"
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	req := authDomain.ReqRefreshToken{
		Agent:        ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		RefreshToken: ctx.MustGet("refreshToken").(string),
	}

	claims := ctx.MustGet("credentials").(*token.AuthClaims)
	res, err := h.authService.RefreshToken(ctx.Request.Context(), claims.Passport, req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
