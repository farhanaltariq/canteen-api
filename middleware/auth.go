package middleware

import (
	"net/http"
	"strings"

	"a21hc3NpZ25tZW50/model"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("error unauthorized user id"))
			return
		}

		authToken := strings.Split(auth, " ")[1]

		if authToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("error unauthorized user id"))
			return
		}

		claims := &model.Claims{}

		tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(err.Error()))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
			return
		}

		if !tkn.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.Set("id", claims.UserID)
		// fmt.Println("\033[32m", "\n\nUSERID\t\n", claims.UserID, "\033[0m")

		ctx.Next()
	})
}
