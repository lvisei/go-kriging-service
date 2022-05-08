//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/lvisei/go-kriging-service/internal/app/api"
	bll "github.com/lvisei/go-kriging-service/internal/app/bll"

	"github.com/google/wire"
	"github.com/lvisei/go-kriging-service/internal/app/router"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		InitGinEngine,
		bll.BllSet,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
