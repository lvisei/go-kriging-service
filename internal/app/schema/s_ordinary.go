package schema

import (
	"github.com/lvisei/go-kriging/ordinarykriging"
)

// Ordinary
//type Ordinary struct {
//	ID        string    `json:"-"` // 唯一标识
//	CreatedAt time.Time `json:"-"` // 请求时间
//	UpdatedAt time.Time `json:"-"` // 返回时间
//}

// OrdinaryTrainParam 训练模型参数
type OrdinaryTrainParam struct {
	Values    []float64                 `json:"values" binding:"required" example:"9.6,10.2,5.4"`      // 权重值数组
	Lons      []float64                 `json:"lons" binding:"required" example:"102.68,99.36,101.23"` // 经度数组
	Lats      []float64                 `json:"lats" binding:"required" example:"25.95,25.81,25.95"`   // 纬度数组
	Sigma2    float64                   `json:"sigma2" binding:"" default:"0" example:"0"`             // sigma2
	Alpha     float64                   `json:"alpha" binding:""  default:"0" example:"100"`           // alpha
	Model     uint16                    `json:"model" binding:"required,oneof=1 2 3" example:"1"`      // 函数模型(1:"spherical" 2:"exponential" 3:"gaussian")
	ModelType ordinarykriging.ModelType `json:"-"`
}

// OrdinaryGridParam 插值网格参数
type OrdinaryGridParam struct {
	Polygon         string                          `json:"polygon" binding:"required" example:"{\"type\": \"Polygon\",\"coordinates\": [[[103.614373, 27.00541],[104.174357, 26.635252],[104.356163, 28.018448],[103.614373, 27.00541]]]}"` // Polygon Geometry String
	PolygonGeometry ordinarykriging.PolygonGeometry `json:"-"`                                                                                                                                                                               // Polygon Geometry
	Width           float64                         `json:"width" binding:"required" example:"0.01" example:"0.01"`                                                                                                                          // 网格单元宽度
}

// OrdinaryPlotPngParam  插值图片参数
type OrdinaryPlotPngParam struct {
	Width  int              `json:"width" binding:"required,gte=1" example:"100"`            // 图片宽度
	Height int              `json:"height" binding:"required,gte=1" example:"100"`           // 图片高度
	Xlim   [2]float64       `json:"xlim" binding:"required" example:"103.614373,104.356163"` // Xlim
	Ylim   [2]float64       `json:"Ylim" binding:"required" example:"26.635252,28.018448"`   // Ylim
	Colors []GridLevelColor `json:"colors" binding:"required,dive,required"`                 // colors
}

type GridLevelColor struct {
	Value [2]float64 `json:"value" binding:"required,unique"  example:"0,15"`             // 值区间 [0, 15]
	Color [4]uint8   `json:"rgba" binding:"dive,gte=0,lte=255" example:"255,128,169,255"` // RGBA颜色 [255, 255, 255, 255]
}

// OrdinaryQueryGridParam 生成插值网格参数
type OrdinaryQueryGridParam struct {
	TrainParam OrdinaryTrainParam `json:"train"`
	GridParam  OrdinaryGridParam  `json:"grid"`
}

// OrdinaryQueryGridPngParam 生成插值图片参数
type OrdinaryQueryGridPngParam struct {
	TrainParam   OrdinaryTrainParam   `json:"train"`
	GridParam    OrdinaryGridParam    `json:"grid"`
	PlotPngParam OrdinaryPlotPngParam `json:"plotPng"`
}

// OrdinaryGridInfo 插值的网格数据
type OrdinaryGridInfo struct {
	GridMatrices *ordinarykriging.GridMatrices `json:"grid"`
	Variogram    *ordinarykriging.Variogram    `json:"-"`
	TimeCost     string                        `json:"timeCost"` // 耗时
}

// OrdinaryQueryOptions 示例对象查询可选参数项
//type OrdinaryQueryOptions struct {
//	OrderFields []*OrderField // 排序字段
//}

// OrdinaryQueryResult 示例对象查询结果
//type OrdinaryQueryResult struct {
//	Data       []*Ordinary
//	PageResult *PaginationResult
//}
