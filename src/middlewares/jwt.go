package middlewares

import (
	"net/http"
	"strings"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/tokens"
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyUserID = "user_id"
	ContextKeyEmail  = "email"
	ContextKeyRoleID = "role_id"
)

// JWTAuth returns a Gin middleware that validates JWT tokens.
func JWTAuth(cfg config.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "authorization header is required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid authorization format, expected: Bearer <token>",
			})
			return
		}

		tokenString := parts[1]
		jwtToken, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](
			tokenString,
			[]byte(cfg.JWT.JWTTokenKey),
			entity.TokenLoginDashboardClaims{},
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid or expired token",
			})
			return
		}

		claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "failed to parse token claims",
			})
			return
		}

		// Set user info in context for downstream handlers
		ctx.Set(ContextKeyUserID, claims.UID)
		ctx.Set(ContextKeyEmail, claims.Email)
		ctx.Set(ContextKeyRoleID, claims.RoleID)

		ctx.Next()
	}
}
