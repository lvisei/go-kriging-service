/*
Package app 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --parseDependency --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger */
package app

// @title go-kriging-service
// @version 1.0.0
// @description GIN + WIRE.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @contact.name
// @contact.email
