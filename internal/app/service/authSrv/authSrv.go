package authSrv

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
	"golang.org/x/crypto/bcrypt"
)

// / When user login remove token that expires
// / Check max device
// When user change the password. Revoke all old token or generate new token

type authSrv struct {
	cfg config.AppConfig

	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(cfg config.AppConfig, authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) authDomain.IAuthSrv {
	return &authSrv{cfg, authRepo, userRepo}
}

func (s *authSrv) Register(ctx context.Context, req authDomain.ReqRegister) error {
	exists, err := s.userRepo.Exists("username", req.Username)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	if exists {
		return ernos.HasAlreadyExists("username")
	}

	req.Password, err = helper.HashPassword(req.Password)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	err = s.authRepo.Register(ctx, req)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (s *authSrv) Login(ctx context.Context, req authDomain.ReqLogin) (res *authDomain.ResToken, err error) {
	var user userDomain.User
	err = s.userRepo.FirstByCol(ctx, &user, "username", req.Username)
	if err != nil {
		return nil, ernos.InvalidCredentials()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, ernos.InvalidCredentials()
	}

	err = s.authRepo.RemoveTokenExpires(ctx, user.ID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	count, err := s.authRepo.Count(ctx, user.ID)
	if count >= int64(s.cfg.Auth.MaxDevice) {
		return nil, ernos.Other(errorDomain.Error{
			Status:  http.StatusForbidden,
			Message: authConst.Msg.MAX_DEVICES_LOGIN,
			Code:    authConst.Code.MAX_DEVICES_LOGIN,
		})
	}

	claims := &authDomain.Passport{
		User: userDomain.User{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
		},
	}

	tokenID := uuid.New()
	newToken := token.New(s.cfg, tokenID, user.Password, *claims)
	_, accessToken, err := newToken.Signed(token.ACCESS)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	refreshClaims, refreshToken, err := newToken.Signed(token.REFRESH)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	res = &authDomain.ResToken{
		ID:           tokenID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	refreshExp, err := time.Parse(time.RFC3339, refreshClaims.ExpiresAt.Format(time.RFC3339))
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	s.authRepo.Create(ctx, authDomain.ReqAuth{
		ID:        tokenID,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserID:    user.ID,
		ExpiresAt: refreshExp,
	})

	return res, nil
}

func (s *authSrv) Logout(ctx context.Context, tokenID uuid.UUID) error {
	return nil
}
