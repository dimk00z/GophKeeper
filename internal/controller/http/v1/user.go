package v1

import (
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) getUserFromCtx(ctx *gin.Context) (user entity.User, err error) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		err = errs.ErrUnexpectedError

		return
	}

	return currentUser.(entity.User), nil
}

func (r *GophKeeperRoutes) UserInfo(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	ctx.JSON(http.StatusOK, currentUser)
}

func (r *GophKeeperRoutes) GetLogins(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	userLogins, err := r.uc.GetLogins(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_logins": userLogins})
}

func (r *GophKeeperRoutes) GetSecretNotes(ctx *gin.Context) {
	userSecretNotes := make([]entity.SecretNote, 0)

	ctx.JSON(http.StatusOK, gin.H{"user_notes": userSecretNotes})
}
