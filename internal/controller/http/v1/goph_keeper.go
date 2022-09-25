package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dimk00z/GophKeeper/internal/usecase"
	"github.com/dimk00z/GophKeeper/pkg/logger"
)

type GophKeeperRoutes struct {
	g usecase.GophKeeper
	l logger.Interface
}

func newGophKeeperRoutes(handler *gin.RouterGroup, g usecase.GophKeeper, l logger.Interface) {
	r := &GophKeeperRoutes{g, l}

	handler.GET("/health", func(ctx *gin.Context) {
		err := g.HealthCheck()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})

			return
		}
		message := "Connection established"
		ctx.JSON(http.StatusOK, gin.H{"status": "connected", "message": message})
	})

	// h := handler.Group("/GophKeeper")
	// {
	// 	h.GET("/history", r.history)
	// 	h.POST("/do-translate", r.doTranslate)
	// }

	userAPI := handler.Group("/user")
	{
		userAPI.GET("me", func(c *gin.Context) {
			c.JSON(http.StatusCreated, struct {
				Response string `json:"response"`
			}{
				Response: "user_data",
			})
		})
	}

	authAPI := handler.Group("/auth")
	{
		authAPI.POST("/register", r.SignUpUser)
		authAPI.POST("/login", r.SignInUser)
		authAPI.GET("/refresh", func(c *gin.Context) {
			c.JSON(http.StatusCreated, struct {
				Response string `json:"response"`
			}{
				Response: "ok",
			})
		})
		authAPI.GET("/logout", func(c *gin.Context) {
			c.JSON(http.StatusCreated, struct {
				Response string `json:"response"`
			}{
				Response: "ok",
			})
		})
	}
}

// type historyResponse struct {
// 	History []entity.GophKeeper `json:"history"`
// }

// @Summary     Show history
// @Description Show all GophKeeper history
// @ID          history
// @Tags  	    GophKeeper
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /GophKeeper/history [get]
// func (r *GophKeeperRoutes) history(c *gin.Context) {
// 	GophKeepers, err := r.g.History(c.Request.Context())
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - history")
// 		errorResponse(c, http.StatusInternalServerError, "database problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, historyResponse{GophKeepers})
// }

// type doTranslateRequest struct {
// 	Source      string `json:"source"       binding:"required"  example:"auto"`
// 	Destination string `json:"destination"  binding:"required"  example:"en"`
// 	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
// }

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
// func (r *GophKeeperRoutes) doTranslate(c *gin.Context) {
// 	var request doTranslateRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	GophKeeper, err := r.g.Translate(
// 		c.Request.Context(),
// 		entity.GophKeeper{
// 			Source:      request.Source,
// 			Destination: request.Destination,
// 			Original:    request.Original,
// 		},
// 	)
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusInternalServerError, "GophKeeper service problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, GophKeeper)
// }
