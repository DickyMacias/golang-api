package handlers

import (
	"movie-tracker/models"
	"movie-tracker/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

func (h *AuthHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (h *AuthHandler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := h.authService.Login(username, password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"title": "Login",
			"error": err.Error(),
			"username": username,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *AuthHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := h.authService.Register(username, email, password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"title": "Register",
			"error": err.Error(),
			"username": username,
			"email": email,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

func (h *AuthHandler) Dashboard(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard",
		"user":  userModel,
	})
}