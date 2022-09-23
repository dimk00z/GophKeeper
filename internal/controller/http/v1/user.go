package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) SignUpUser(ctx *gin.Context) {
	var payload *struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	user, err := r.g.SignUpUser(ctx, payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, user)

}
