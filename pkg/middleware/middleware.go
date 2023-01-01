package middleware

import (
	"github.com/labstack/echo/v4"
)

var SortBy *string

const (
	limitConst         = "5"
	defaultOffsetParam = "0"
)

func PaginationAdvertisement(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("limit", limitConst)
		c.Request().Header.Set("offset", defaultOffsetParam)

		err := next(c)
		if err != nil {
			return err
		}
		return nil
	}
}
