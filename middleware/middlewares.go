package middleware

import (
	"Go-Learn-API/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware := Validate jwt token of incoming request
func TokenAuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		err := auth.ValidateToken(ctx.Request)	
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "you need to be authorized to access this route")
			ctx.Abort()
			return 
		}
		ctx.Next()
	}	
}