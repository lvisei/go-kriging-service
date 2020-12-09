package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrdinarySet 注入Ordinary
var OrdinarySet = wire.NewSet(wire.Struct(new(Ordinary), "*"))

// Ordinary 普通克力金
type Ordinary struct {
}

// Query 查询插值生成的网格数据
// @Tags 普通克力金
// @Summary 查询插值生成的网格数据
// @Param body body schema.OrdinaryQueryGridParam true "请求参数"
// @Success 200 {object} schema.OrdinaryGridInfo "插值的网格数据"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/ordinary/grid [post]
func (a *Ordinary) Grid(c *gin.Context) {
}

// Get 查询插值生成的图片
// @Tags 普通克力金
// @Summary 查询插值生成图片
// @Produce png
// @Param body body schema.OrdinaryQueryGridPngParam true "请求参数"
// @Success 200 "文件图片"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/ordinary/grid-png [post]
func (a *Ordinary) GridPng(c *gin.Context) {
}
