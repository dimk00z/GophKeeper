package v1

import (
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) GetCards(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userCards, err := r.uc.GetCards(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userCards) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"user_cards": userCards})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_cards": userCards})
}

func (r *GophKeeperRoutes) AddCard(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	var payloadCard *entity.Card

	if err := ctx.ShouldBindJSON(&payloadCard); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := r.uc.AddCard(ctx, payloadCard, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusAccepted, payloadCard)
}
