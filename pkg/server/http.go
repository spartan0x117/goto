package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spartan0x117/goto/pkg/storage"
)

func NewServer(s storage.Storage) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	addHandler := func(c echo.Context) error {
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
				return c.String(http.StatusNotFound, "no link found for label")
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

	openHandler := func(c echo.Context) error {
		label := c.Param("label")
		if label == "" {
			return c.String(http.StatusOK, "please enter a label...")
		}
		label, path, pathExists := strings.Cut(label, "/")
		link := s.GetLinkForLabel(label)

		if link == "" {
			return c.String(http.StatusNotFound, "label not found")
		}
		if pathExists {
			link, _ = url.JoinPath(link, path)
		}

		return c.Redirect(http.StatusFound, link)
	}
	e.GET("/:label", openHandler)

	removeHandler := func(c echo.Context) error {
		if c.Param("label") == "" {
			return c.String(http.StatusBadRequest, "please include a label to delete...")
		}

		err := s.RemoveLink(c.Param("label"))
		if err != nil {
			return c.String(http.StatusInternalServerError, "error trying to delete label...")
		}
		return c.String(http.StatusAccepted, "removed label")
	}
	e.GET("/remove/:label", removeHandler)

	syncHandler := func(c echo.Context) error {
		err := s.Sync()
		if err != nil {
			return c.String(http.StatusInternalServerError, "error trying to sync...")
		}
		return c.String(http.StatusOK, "done syncing")
	}
	e.GET("/sync", syncHandler)
	e.GET("/sync/", syncHandler)

	return e
}
