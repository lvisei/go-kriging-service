package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lvisei/go-kriging-service/internal/app/bll"
	"github.com/lvisei/go-kriging-service/internal/app/ginx"
	"github.com/lvisei/go-kriging-service/internal/app/schema"
	"github.com/lvisei/go-kriging-service/pkg/errors"
	"github.com/lvisei/go-kriging-service/pkg/util/json"
	"github.com/lvisei/go-kriging/ordinarykriging"
)

// OrdinarySet 注入Ordinary
var OrdinarySet = wire.NewSet(wire.Struct(new(Ordinary), "*"))

// Ordinary
type Ordinary struct {
	OrdinaryBll *bll.Ordinary
}

// Grid 插值生成网格数据
func (a *Ordinary) Grid(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrdinaryQueryGridParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	if err := json.Unmarshal([]byte(item.GridParam.Polygon), &item.GridParam.PolygonGeometry); err != nil {
		ginx.ResError(c, err)
		return
	}

	gridInfo, err := a.OrdinaryBll.Grid(ctx, item.TrainParam, item.GridParam)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, gridInfo)
}

// GridPng 插值生成网格图片
func (a *Ordinary) GridPng(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrdinaryQueryGridPngParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	var polygonGeometry ordinarykriging.PolygonGeometry
	if err := json.Unmarshal([]byte(item.GridParam.Polygon), &polygonGeometry); err != nil {
		ginx.ResError(c, err)
		return
	} else if polygonGeometry.Coordinates == nil || len(polygonGeometry.Coordinates) == 0 || len(polygonGeometry.Coordinates[0]) == 0 {
		// TODO: validator validate polygonGeometry
		ginx.ResError(c, errors.New400Response("Polygon 类型错误"))
		return
	}

	item.GridParam.PolygonGeometry = polygonGeometry
	gridInfo, err := a.OrdinaryBll.Grid(ctx, item.TrainParam, item.GridParam)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	err = a.OrdinaryBll.GridPng(ctx, c.Writer, gridInfo, item.PlotPngParam)

	if err != nil {
		ginx.ResError(c, err)
		return
	}
}
