package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) ProtectedByAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		accessTokenFromCookie, err := ctx.Cookie("access_token")

		authorizationHeader := strings.Fields(ctx.Request.Header.Get("Authorization"))

		if len(authorizationHeader) != 0 && authorizationHeader[0] == "Bearer" {
			accessToken = authorizationHeader[1]
		} else if err == nil {
			accessToken = accessTokenFromCookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		user, err := r.uc.CheckAccessToken(ctx, accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
