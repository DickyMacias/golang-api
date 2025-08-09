package routes

import (
	"movie-tracker/handlers"
	"movie-tracker/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.SessionMiddleware())

	authHandler := handlers.NewAuthHandler()
	tmdbHandler := handlers.NewTMDBHandler()
	favoritesHandler := handlers.NewFavoritesHandler()
	userHandler := handlers.NewUserHandler()

	// Root redirect
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

	// Test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{
			"message": "Hello World",
		})
	})

	// Public routes (redirect if authenticated)
	public := r.Group("/")
	public.Use(middleware.RedirectIfAuthenticated())
	{
		public.GET("/login", authHandler.ShowLogin)
		public.POST("/login", authHandler.Login)
		public.GET("/register", authHandler.ShowRegister)
		public.POST("/register", authHandler.Register)
	}

	// Authentication logout (available to authenticated users)
	r.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Dashboard
		protected.GET("/dashboard", authHandler.Dashboard)

		// Movie pages
		protected.GET("/search", favoritesHandler.ShowSearch)
		protected.GET("/favorites", favoritesHandler.ShowFavorites)
		protected.GET("/movie/:id", tmdbHandler.GetMovieDetail)
		protected.GET("/favorites/por-ver", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/favorites?status=por_ver")
		})
		protected.GET("/favorites/vistas", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/favorites?status=vista")
		})
		protected.GET("/favorites/recomendadas", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/favorites?status=recomendada")
		})
	}

	// API routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// TMDB API
		api.GET("/movies/search", tmdbHandler.SearchMovies)
		api.GET("/movies/popular", tmdbHandler.GetPopularMovies)
		api.GET("/movies/trending", tmdbHandler.GetTrendingMovies)

		// Favorites API
		api.POST("/favorites", favoritesHandler.AddToFavorites)
		api.PATCH("/favorites/:id/status", favoritesHandler.UpdateStatus)
		api.PATCH("/favorites/:id/rating", favoritesHandler.UpdateRating)
		api.DELETE("/favorites/:id", favoritesHandler.DeleteFavorite)

		// Stats API
		api.GET("/stats", userHandler.GetStats)
	}
}