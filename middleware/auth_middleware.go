package middleware

import (
	"movie-tracker/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		authService := services.NewAuthService()
		user, err := authService.GetUserByID(userID.(uint))
		if err != nil {
			session.Delete("user_id")
			session.Save()
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func RedirectIfAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID != nil {
			c.Redirect(http.StatusFound, "/dashboard")
			c.Abort()
			return
		}

		c.Next()
	}
}