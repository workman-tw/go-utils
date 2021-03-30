package echo

import (
	"github.com/labstack/echo/v4"
)

// Route - define router
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// RouteGroups - define router groups
type RouteGroups struct {
	Prefix     string
	Routes     []*Route
	Mideleware []echo.MiddlewareFunc
}
