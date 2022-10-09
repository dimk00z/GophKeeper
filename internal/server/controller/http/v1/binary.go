package v1

import (
	"errors"
	"net/http"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

var errBinaryNameNotGiven = errors.New("binary name has not given")

// @Summary     Get user binary data
// @Description fetching user binary data
// @ID          get_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Binary
// @Success     204
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary [get].
func (r *GophKeeperRoutes) GetBinaries(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userBinaries, err := r.uc.GetBinaries(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userBinaries) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userBinaries)
}

// @Summary     Add user binary data
// @Description saving user binary data
// @ID          add_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Binary
// @Success     204
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary [post].
func (r *GophKeeperRoutes) AddBinary(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}
	var binary entity.Binary
	if ctx.Query("name") == "" {
		errorResponse(ctx, http.StatusBadRequest, errBinaryNameNotGiven.Error())

		return
	}
	binary.Name = ctx.Query("name")

	file, err := ctx.FormFile("file")
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	binary.FileName = file.Filename
	if err = r.uc.AddBinary(ctx, &binary, file, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	ctx.JSON(http.StatusOK, binary)
}
