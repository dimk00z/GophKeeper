package v1

import (
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) UserInfo(ctx *gin.Context) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrUnexpectedError.Error()})
	}

	ctx.JSON(http.StatusOK, currentUser.(entity.User))
}

func (r *GophKeeperRoutes) GetLogins(ctx *gin.Context) {
	userLogins := make([]entity.Login, 0)

	ctx.JSON(http.StatusOK, gin.H{"user_logins": userLogins})
}

func (r *GophKeeperRoutes) GetCards(ctx *gin.Context) {
	userCards := make([]entity.Card, 0)

	ctx.JSON(http.StatusOK, gin.H{"user_cards": userCards})
}

func (r *GophKeeperRoutes) GetSecretNotes(ctx *gin.Context) {
	userSecretNotes := make([]entity.SecretNote, 0)

	ctx.JSON(http.StatusOK, gin.H{"user_notes": userSecretNotes})
}
