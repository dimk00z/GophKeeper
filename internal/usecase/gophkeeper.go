package usecase

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	config "github.com/dimk00z/GophKeeper/config/server"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/patrickmn/go-cache"
)

const minutesPerHour = 60

// GophKeeperUseCase -.
type GophKeeperUseCase struct {
	repo  GophKeeperRepo
	cfg   *config.Config
	cache *cache.Cache
}

func New(r GophKeeperRepo, cfg *config.Config) *GophKeeperUseCase {
	return &GophKeeperUseCase{
		repo: r,
		cfg:  cfg,
		cache: cache.New(
			time.Duration(cfg.Cache.DefaultExpiration)*time.Minute,
			time.Duration(cfg.Cache.CleanupInterval)*time.Minute),
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

func (uc *GophKeeperUseCase) CheckAccessToken(ctx context.Context, accessToken string) (entity.User, error) {
	if userFromCache, found := uc.cache.Get(accessToken); found {
		checkedUser, ok := userFromCache.(entity.User)

		if ok {
			return checkedUser, nil
		}
	}

	var user entity.User

	sub, err := utils.ValidateToken(accessToken, uc.cfg.AccessTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return user, err
	}

	userID := fmt.Sprint(sub)
	user, err = uc.repo.GetUserByID(ctx, userID)

	if err != nil {
		err = errs.ErrTokenValidation

		return user, err
	}

	uc.cache.Set(accessToken, user, cache.DefaultExpiration)

	return user, nil
}
