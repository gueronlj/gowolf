package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render takes a context and a template component. It can return an error
func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
