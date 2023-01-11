package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spartan0x117/goto/pkg/storage"
)

func NewServer(s storage.Storage) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	addHandler := func(c echo.Context) error {
		e.Logger.Info("inside 'addHandler'")
		if c.Request().Method == http.MethodPost {
			label, link := c.FormValue("label"), c.FormValue("url")
			err := s.AddLink(label, link, true) // TODO: Some kind of confirmation of overwriting a link
			if err != nil {
				return c.String(http.StatusBadRequest, "could not add goto link...")
			}
			return c.String(http.StatusAccepted, fmt.Sprintf("added '%s: %s'", label, link))
		}
		return c.HTML(http.StatusOK, addFormHtml)
	}
	e.GET("/add", addHandler)
	e.GET("/add/", addHandler)
	e.POST("/add", addHandler)
	e.POST("/add/", addHandler)

	findHandler := func(c echo.Context) error {
		label := c.Param("label")
		if label != "" {
			link := s.GetLinkForLabel(label)
			if link == "" {
				return c.String(http.StatusNotFound, fmt.Sprintf("no link found for '%s'", label))
			}
			return c.String(http.StatusOK, link)
		}

		allLinks := s.GetAllLabels()
		var sb strings.Builder
		for _, link := range allLinks {
			sb.WriteString(fmt.Sprintf("%s\n", link))
		}
		return c.String(http.StatusOK, sb.String())
	}
	e.GET("/find", findHandler)
	e.GET("/find/", findHandler)
	e.GET("/find/:label", findHandler)

	gotoHandler := func(c echo.Context) error {
		e.Logger.Info("inside 'gotoHandler'")
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
