package middleware

import (
	"github.com/labstack/echo/v4"
)

const (
	ascendingDateSort = "desc"
	descendingDateSort = ""
)

func SortAdvByDate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		
		err := next(c)
		if err != nil {
			return err
		}
		return nil
	}
}

