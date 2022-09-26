package v1

import (
	"errors"
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *GophKeeperRoutes) SignUpUser(ctx *gin.Context) {
	var payload *loginPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	user, err := r.uc.SignUpUser(ctx, payload.Email, payload.Password)
	if err == nil {
		ctx.JSON(http.StatusCreated, user)

		return
	}

	if errors.Is(err, errs.ErrWrongEmail) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *GophKeeperRoutes) SignInUser(ctx *gin.Context) {
	var payload *loginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	jwt, err := r.uc.SignInUser(ctx, payload.Email, payload.Password)

	if err == nil {
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("refresh_token", jwt.RefreshToken, jwt.RefreshTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

		ctx.JSON(http.StatusOK, jwt)

		return
	}

	if errors.Is(err, errs.ErrWrongCredentials) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *GophKeeperRoutes) RefreshAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "refresh token has not been found"})

		return
	}

	jwt, err := r.uc.RefreshAccessToken(ctx, refreshToken)

	if err == nil {
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

		ctx.JSON(http.StatusOK, jwt)

		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *GophKeeperRoutes) LogoutUser(ctx *gin.Context) {
	domainName := r.uc.GetDomainName()
	ctx.SetCookie("access_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("logged_in", "", -1, "/", domainName, false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (r *GophKeeperRoutes) UserInfo(ctx *gin.Context) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrUnexpectedError.Error()})
	}

	ctx.JSON(http.StatusOK, currentUser.(entity.User))
}
