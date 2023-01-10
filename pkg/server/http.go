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
	//e.Use(middleware.Logger())

	addHandler := func(c echo.Context) error {
		e.Logger.Info("inside 'addHandler'")
		if c.Request().Method == http.MethodPost {
			label, link := c.FormValue("label"), c.FormValue("url")
			err := s.AddLink(label, link)
			if err != nil {
				return c.String(http.StatusBadRequest, "could not add goto link...")
			}
			return c.String(http.StatusAccepted, fmt.Sprintf("added '%s: %s'", label, link))
		}
		return c.File("./pkg/server/add_form.html")
	}
	e.GET("/add", addHandler)
	e.POST("/add", addHandler)

	gotoHandler := func(c echo.Context) error {
		e.Logger.Info("inside 'gotoHandler'")
		label := c.Param("label")
		if label == "" {
			return c.String(http.StatusOK, "Please enter a label: go/<label>")
		}
		link := s.GetLinkForLabel(label)
		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			link = "https://" + link
		}
		u, err := url.Parse(link)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing url: %s", link))
		}
		if link == "" {
			return c.String(http.StatusNotFound, fmt.Sprintf("Label '%s' not found", label))
		}

		return c.Redirect(http.StatusFound, u.String())
	}

	e.GET("/:label", gotoHandler)

	return e
}
