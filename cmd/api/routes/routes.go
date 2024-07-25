package routes

import (
	middleware "pantori/cmd/api/middlewares"
	"pantori/internal/auth"
	"pantori/internal/domains/categories"
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
	middlewares := middleware.New()
	authRoutes := auth.NewNetworkHandler()
	goodsRoutes := goods.New()
	categoriesRoutes := categories.New()

	router := gin.New()

	router.Use(middlewares.SetCORS())
	router.Use(middlewares.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := router.Group("/api")
	{
		api.POST("/login", authRoutes.Login)
		auth := api.Group("/auth", middlewares.AuthorizeRequest(middleware.AuthorizeInput{AdminRequired: true}))
		{
			auth.POST("/user", authRoutes.CreateUser)
			auth.DELETE("/user", authRoutes.DeleteUser)
			auth.GET("/user", authRoutes.ListUsers)
		}

		goods := api.Group("/goods", middlewares.AuthorizeRequest(middleware.AuthorizeInput{AdminRequired: false}))
		{
			goods.POST("", goodsRoutes.CreateGood)
			goods.PATCH("", goodsRoutes.EditGood)
			goods.GET("", goodsRoutes.ListGoods)
			goods.GET("/:id", goodsRoutes.GetGood)
			goods.DELETE("", goodsRoutes.DeleteGood)
			goods.GET("/shopping-list", goodsRoutes.GetShoppingList)
		}

		categories := api.Group("/categories", middlewares.AuthorizeRequest(middleware.AuthorizeInput{AdminRequired: false}))
		{
			categories.GET("", categoriesRoutes.ListCategories)
			categories.POST("", categoriesRoutes.CreateCategory)
			categories.POST("/default", categoriesRoutes.CreateDefaultCategories)
			categories.DELETE("", categoriesRoutes.DeleteCategory)
			categories.PATCH("", categoriesRoutes.EditCategory)
		}
	}
	router.Run(":8800")
}
