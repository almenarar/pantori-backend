package goodshdl

import (
	core "pantori/internal/domains/goods/core"

	"net/http"

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
// @Summary Register a good
// @Description Endpoint used to Create a single good in database
// @Tags Goods
// @Accept json
// @Produce json
// @Param PostGood body goodscore.PostGood true "PostGood"
// @Success 200 {string} ok
// @Router /goods [post]
// @Security ApiKeyAuth
func (net *Network) CreateGood(ctx *gin.Context) {
	payload := core.PostGood{}
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

	if err := net.svc.AddGood(
		core.Good{
			Name:      payload.Name,
			Category:  payload.Category,
			Workspace: payload.Workspace,
			Expire:    payload.Expire,
			BuyDate:   payload.BuyDate,
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
// @Summary List goods
// @Description Endpoint used to List all goods from a workspace in database
// @Tags Goods
// @Produce json
// @Success 200 {string} arn
// @Router /goods [get]
// @Security ApiKeyAuth
func (net *Network) ListGoods(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if exists && username == "dryrun" {
		ctx.JSON(
			http.StatusOK,
			[]core.Good{
				{
					Name:     "dryrun1",
					Category: "categoryDR1",
					BuyDate:  "2015/03/01",
					Expire:   "2015/03/01",
				},
				{
					Name:     "dryrun2",
					Category: "categoryDR2",
					BuyDate:  "2015/03/01",
					Expire:   "2015/03/01",
				},
			},
		)
		return
	}

	output, err := net.svc.ListGoods()
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

// PingExample godoc
// @Summary Delete a good
// @Description Endpoint used to Delete a single good from database by ID
// @Tags Goods
// @Accept json
// @Produce json
// @Param DeleteGood body goodscore.DeleteGood true "DeleteGood"
// @Success 200 {string} ok
// @Router /goods [delete]
// @Security ApiKeyAuth
func (net *Network) DeleteGood(ctx *gin.Context) {
	payload := core.DeleteGood{}
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

	if err := net.svc.DeleteGood(
		core.Good{
			ID: payload.ID,
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
