package main

import (
	"net/http"
	"os"

	"github.com/syukur91/ischool/student"
	"github.com/syukur91/ischool/studentserver"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func setupHandlers(e *echo.Echo, log *log.Entry) {

	r := e.Group("/:tenant/api")

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello "+c.Param("tenant")+"! This is API version: "+os.Getenv("VERSION"))
	})

	studentClient := student.NewStudentCLI(log)

	studentserverHandler := &studentserver.Handler{
		StudentClient: studentClient,
		Log:           log,
	}

	studentserverHandler.SetRoutes(r)
}
