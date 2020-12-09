package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuvigongzuoshi/go-kriging-service/internal/app/middleware"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{

		gOrdinary := v1.Group("ordinary")
		{
			gOrdinary.POST("grid", a.OrdinaryAPI.Grid)
			gOrdinary.POST("grid-png", a.OrdinaryAPI.GridPng)
		}

	}
}
