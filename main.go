package main

import (
	"product-api/app/handler"
	"product-api/app/middleware"
	"product-api/app/repository"
	"product-api/app/use_case"
	"product-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	repoProduct := repository.NewProductRepository(db)
	ucProduct := use_case.NewProductUseCase(repoProduct)

	e := echo.New()
	handler.NewProductHandler(e, ucProduct)
	middleware.LoggerMiddleware(e)
	
	e.Logger.Fatal(e.Start(":8000"))
}