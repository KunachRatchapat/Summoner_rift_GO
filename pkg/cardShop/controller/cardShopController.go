package controller

import "github.com/labstack/echo/v4"

type CardShopController interface {
	Listing(pctx echo.Context ) error
}