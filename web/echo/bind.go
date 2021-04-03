package echo

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/workman-tw/go-utils/web"
)

// BindAndValidate - binding data and validate
func BindAndValidate(c echo.Context, postData interface{}) error {
	if err := c.Bind(postData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, web.BindError(err).Error())
	}

	if err := c.Validate(postData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, web.ValidateError(err).Error())
	}

	return nil
}
