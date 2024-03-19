package routes

import (
	"pantori/cmd/api/middlewares"
	"pantori/internal/auth"
	"pantori/internal/domains/goods"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type routes struct{}

func New() *routes {
	return &routes{}
}

func (r *routes) Expose() {
	middlewares := middlewares.New()
	auth := auth.New()
	goodsRoutes := goods.New()

	router := gin.New()

	router.Use(middlewares.SetCORS())
	router.Use(middlewares.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := router.Group("/api")
	{
		api.POST("/login", auth.Login)

		goods := api.Group("/goods", middlewares.AuthorizeRequest())
		{
			goods.POST("", goodsRoutes.CreateGood)
			goods.PATCH("", goodsRoutes.EditGood)
			goods.GET("", goodsRoutes.ListGoods)
			goods.DELETE("", goodsRoutes.DeleteGood)
		}
	}
	router.Run(":8800")
}
