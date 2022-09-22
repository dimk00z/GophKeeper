package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase"
	"github.com/dimk00z/GophKeeper/pkg/logger"
)

type GophKeeperRoutes struct {
	t usecase.GophKeeper
	l logger.Interface
}

func newGophKeeperRoutes(handler *gin.RouterGroup, t usecase.GophKeeper, l logger.Interface) {
	r := &GophKeeperRoutes{t, l}

	h := handler.Group("/GophKeeper")
	{
		h.GET("/history", r.history)
		h.POST("/do-translate", r.doTranslate)
	}
}

type historyResponse struct {
	History []entity.GophKeeper `json:"history"`
}

// @Summary     Show history
// @Description Show all GophKeeper history
// @ID          history
// @Tags  	    GophKeeper
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /GophKeeper/history [get]
func (r *GophKeeperRoutes) history(c *gin.Context) {
	GophKeepers, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{GophKeepers})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    GophKeeper
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up GophKeeper"
// @Success     200 {object} entity.GophKeeper
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /GophKeeper/do-translate [post]
func (r *GophKeeperRoutes) doTranslate(c *gin.Context) {
	var request doTranslateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	GophKeeper, err := r.t.Translate(
		c.Request.Context(),
		entity.GophKeeper{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "GophKeeper service problems")

		return
	}

	c.JSON(http.StatusOK, GophKeeper)
}
