package handlers

import (
	"net/http"
	"pantori/internal/domains/categories/core"

	"github.com/gin-gonic/gin"
)

type Network struct {
	svc core.ServicePort
}

func NewNetwork(svc core.ServicePort) *Network {
	return &Network{
		svc: svc,
	}
}

// PingExample godoc
// @Summary Register a category
// @Description Endpoint used to Create a single category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param PostCategory body core.PostCategory true "PostCategory"
// @Success 200 {string} ok
// @Router /categories [post]
// @Security ApiKeyAuth
func (net *Network) CreateCategory(ctx *gin.Context) {
	payload := core.PostCategory{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "some of the required fields are empty",
			},
		)
		return
	}

	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			"Dry run execution ok",
		)
		return
	}

	if err := net.svc.CreateCategory(
		core.Category{
			Name:      payload.Name,
			Workspace: payload.Workspace,
			Color:     payload.Color,
		},
	); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		"Done",
	)
}

// PingExample godoc
// @Summary Register default categories for a workspace
// @Description Endpoint used to Create default categories in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param PostDefaultCategories body core.PostDefaultCategories true "PostDefaultCategories"
// @Success 200 {string} ok
// @Router /categories/default [post]
// @Security ApiKeyAuth
func (net *Network) CreateDefaultCategories(ctx *gin.Context) {
	payload := core.PostDefaultCategories{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "some of the required fields are empty",
			},
		)
		return
	}

	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			"Dry run execution ok",
		)
		return
	}

	if err := net.svc.CreateDefaultCategories(
		payload.Workspace,
	); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		"Done",
	)
}

// PingExample godoc
// @Summary Delete a category
// @Description Endpoint used to Delete a single category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param DeleteCategory body core.DeleteCategory true "DeleteCategory"
// @Success 200 {string} ok
// @Router /categories [delete]
// @Security ApiKeyAuth
func (net *Network) DeleteCategory(ctx *gin.Context) {
	payload := core.DeleteCategory{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "some of the required fields are empty",
			},
		)
		return
	}

	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			"Dry run execution ok",
		)
		return
	}

	if err := net.svc.DeleteCategory(
		core.Category{
			ID:        payload.ID,
			Workspace: payload.Workspace,
		},
	); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		"Done",
	)
}

// PingExample godoc
// @Summary Edit a category
// @Description Endpoint used to Edit a single category in database
// @Tags Categories
// @Accept json
// @Produce json
// @Param PatchCategory body core.PatchCategory true "PatchCategory"
// @Success 200 {string} ok
// @Router /categories [patch]
// @Security ApiKeyAuth
func (net *Network) EditCategory(ctx *gin.Context) {
	payload := core.PatchCategory{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "some of the required fields are empty",
			},
		)
		return
	}

	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			"Dry run execution ok",
		)
		return
	}

	if err := net.svc.EditCategory(
		core.Category{
			ID:        payload.ID,
			Name:      payload.Name,
			Color:     payload.Color,
			Workspace: payload.Workspace,
		},
	); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		"Done",
	)
}

// PingExample godoc
// @Summary List categories
// @Description Endpoint used to List all categories from a workspace in database
// @Tags Categories
// @Param workspace path string true "Workspace"
// @Produce json
// @Success 200 {string} arn
// @Router /categories/{workspace} [get]
// @Security ApiKeyAuth
func (net *Network) ListCategories(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			[]core.Category{
				{
					Name:  "dryrun1",
					Color: "green",
				},
				{
					Name:  "dryrun2",
					Color: "blue",
				},
			},
		)
		return
	}

	output, err := net.svc.ListCategories(ctx.Param("workspace"))
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		output,
	)
}
