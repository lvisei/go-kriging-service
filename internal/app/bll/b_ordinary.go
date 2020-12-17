package bll

import (
	"context"
	"fmt"
	"github.com/liuvigongzuoshi/go-kriging-service/internal/app/schema"
	"github.com/liuvigongzuoshi/go-kriging-service/pkg/errors"
	"image/color"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/liuvigongzuoshi/go-kriging/ordinarykriging"
)

// OrdinarySet 注入Ordinary
var OrdinarySet = wire.NewSet(wire.Struct(new(Ordinary), "*"))

// Ordinary
type Ordinary struct {
}

// Grid 插值生成网格数据
func (a *Ordinary) Grid(ctx context.Context, trainParam schema.OrdinaryTrainParam, gridParam schema.OrdinaryGridParam) (*schema.OrdinaryGridInfo, error) {
	model := trainParam.Model
	if model == 1 {
		trainParam.ModelType = ordinarykriging.Spherical
	} else if model == 2 {
		trainParam.ModelType = ordinarykriging.Exponential
	} else if model == 3 {
		trainParam.ModelType = ordinarykriging.Gaussian
	} else {
		return nil, errors.New400Response("模型类型错误")
	}

	start := time.Now()

	ordinaryKriging := ordinarykriging.NewOrdinary(trainParam.Values, trainParam.Lons, trainParam.Lats)
	if _, err := ordinaryKriging.Train(trainParam.ModelType, trainParam.Sigma2, trainParam.Alpha); err != nil {
		return nil, errors.New400Response("训练模型参数有误")
	}

	gridMatrices := ordinaryKriging.Grid(gridParam.PolygonGeometry.Coordinates, gridParam.Width)

	tc := time.Since(start)
	timeCost := fmt.Sprintf("time cost = %v s", tc.Seconds())

	gridInfo := &schema.OrdinaryGridInfo{GridMatrices: gridMatrices, Variogram: ordinaryKriging, TimeCost: timeCost}

	return gridInfo, nil
}

// GridPng 插值生成网格图片
func (a *Ordinary) GridPng(ctx context.Context, w http.ResponseWriter, gridInfo *schema.OrdinaryGridInfo, params schema.OrdinaryPlotPngParam) error {
	ordinaryKriging := gridInfo.Variogram
	var gridLevelColor []ordinarykriging.GridLevelColor
	for _, item := range params.Colors {
		gridLevelColor = append(gridLevelColor, ordinarykriging.GridLevelColor{Value: item.Value, Color: color.RGBA{R: item.Color[0], G: item.Color[1], B: item.Color[2], A: item.Color[3]}})
	}
	canvasX := ordinaryKriging.Plot(gridInfo.GridMatrices, params.Width, params.Height, params.Xlim, params.Ylim, gridLevelColor)

	//subTitle := &canvas.TextConfig{
	//	Text:     "球面半变异函数模型",
	//	FontName: "testdata/fonts/source-han-sans-sc/regular.ttf",
	//	FontSize: 28,
	//	Color:    color.RGBA{R: 0, G: 0, B: 0, A: 255},
	//	OffsetX:  252,
	//	OffsetY:  40,
	//	AlignX:   0.5,
	//}
	//if err := canvasX.DrawText(subTitle); err != nil {
	//	return nil, err
	//}

	buffer, err := canvasX.Output()
	if err != nil {
		return err
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	w.Write(buffer)

	//filePath := fmt.Sprintf("%v/%v %v", "tmp", time.Now().Format("2006-01-02 15-04-05"), "grid-png.png")
	//if err := canvasX.SavePNG(filePath); err != nil {
	//	errors.New500Response("保存插值生成网格图片失败")
	//}

	return nil
}
