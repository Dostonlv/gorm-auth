package api

import (
	"github.com/Dostonlv/gorm-auth/api/docs"
	"github.com/Dostonlv/gorm-auth/config"
	"github.com/Dostonlv/gorm-auth/controller"
	"github.com/Dostonlv/gorm-auth/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @description	This is a sample Auth demo.
// @termsOfService	https://Doston.me

func SetUpAPI(r *gin.Engine, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.ServicePort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// SignUP godoc
	// Tags signup
	// @ID signup
	// @Summary Create account
	// @Description Create account
	// @Param data body models.User true "User Body"
	// @Accept  json
	// @Produce  json
	// @Success 200
	// @Failure default
	// @Router /signup [POST]
	r.POST("/signup", controller.SignUp)

	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.MiddlewareAuth, controller.Validate)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
