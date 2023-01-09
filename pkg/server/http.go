package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spartan0x117/goto/pkg/storage"
)

func NewServer(s storage.Storage) *echo.Echo {
	e := echo.New()
	gotoHandler := func(c echo.Context) error {
		label := c.Param("label")
		if label == "" {
			return c.String(http.StatusOK, "Please enter a label: go/<label>")
		}
		link := s.GetLinkForLabel(label)
		if link == "" {
			return c.String(http.StatusNotFound, fmt.Sprintf("Label '%s' not found", label))
		}

		return c.Redirect(http.StatusFound, link)
	}

	e.GET("/:label", gotoHandler)

	return e
}
