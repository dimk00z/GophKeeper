package v1

import (
	"errors"
	"fmt"
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	user, err := r.g.SignUpUser(ctx, payload.Email, payload.Password)
	if err == nil {
		ctx.JSON(http.StatusCreated, user)
		return

	}
	if errors.Is(err, errs.ErrWrongEmail) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

}

func (r *GophKeeperRoutes) SignInUser(ctx *gin.Context) {
	var payload *loginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})

		return
	}
	jwt, err := r.g.SignInUser(ctx, payload.Email, payload.Password)
	fmt.Println(jwt, err)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": jwt.AccessToken})

		return
	}
	// https://codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens/
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

}
