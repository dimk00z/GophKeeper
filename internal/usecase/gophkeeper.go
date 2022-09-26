package usecase

import (
	"context"
	"fmt"
	"net/mail"

	config "github.com/dimk00z/GophKeeper/config/server"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
)

const minutesPerHour = 60

// GophKeeperUseCase -.
type GophKeeperUseCase struct {
	repo   GophKeeperRepo
	webAPI GophKeeperWebAPI
	cfg    *config.Config
}

func New(r GophKeeperRepo, w GophKeeperWebAPI, cfg *config.Config) *GophKeeperUseCase {
	return &GophKeeperUseCase{
		repo:   r,
		webAPI: w,
		cfg:    cfg,
	}
}

func (uc *GophKeeperUseCase) HealthCheck() error {
	return uc.repo.DBHealthCheck()
}

func (uc *GophKeeperUseCase) SignUpUser(ctx context.Context, email, password string) (user entity.User, err error) {
	if _, err = mail.ParseAddress(email); err != nil {
		err = errs.ErrWrongEmail

		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return
	}

	return uc.repo.AddUser(ctx, email, hashedPassword)
}

func (uc *GophKeeperUseCase) SignInUser(ctx context.Context, email, password string) (token entity.JWT, err error) {
	if _, err = mail.ParseAddress(email); err != nil {
		err = errs.ErrWrongEmail

		return
	}

	user, err := uc.repo.GetUserByEmail(ctx, email, password)
	if err != nil {
		return
	}

	token.AccessToken, err = utils.CreateToken(
		uc.cfg.Secutiry.AccessTokenExpiresIn,
		user.ID,
		uc.cfg.Secutiry.AccessTokenPrivateKey)
	if err != nil {
		return
	}

	token.RefreshToken, err = utils.CreateToken(
		uc.cfg.Secutiry.RefreshTokenExpiresIn,
		user.ID,
		uc.cfg.Secutiry.RefreshTokenPrivateKey)

	if err != nil {
		return
	}

	token.AccessTokenMaxAge = uc.cfg.Secutiry.AccessTokenMaxAge * minutesPerHour
	token.RefreshTokenMaxAge = uc.cfg.Secutiry.RefreshTokenMaxAge * minutesPerHour
	token.Domain = uc.cfg.Secutiry.Domain

	return token, nil
}

func (uc *GophKeeperUseCase) RefreshAccessToken(ctx context.Context, refreshToken string) (token entity.JWT, err error) {
	userID, err := utils.ValidateToken(refreshToken, uc.cfg.Secutiry.RefreshTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return
	}

	user, err := uc.repo.GetUserByID(ctx, fmt.Sprint(userID))
	if err != nil {
		err = errs.ErrTokenValidation

		return
	}

	token.RefreshToken = refreshToken
	token.AccessToken, err = utils.CreateToken(
		uc.cfg.Secutiry.AccessTokenExpiresIn,
		user.ID,
		uc.cfg.Secutiry.AccessTokenPrivateKey)

	if err != nil {
		return
	}

	token.AccessTokenMaxAge = uc.cfg.Secutiry.AccessTokenMaxAge * minutesPerHour
	token.RefreshTokenMaxAge = uc.cfg.Secutiry.RefreshTokenMaxAge * minutesPerHour
	token.Domain = uc.cfg.Secutiry.Domain

	return
}

func (uc *GophKeeperUseCase) GetDomainName() string {
	return uc.cfg.Secutiry.Domain
}
