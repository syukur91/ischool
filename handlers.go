package main

import (
	"net/http"
	"os"

	"github.com/syukur91/ischool/domain"
	postgreStorage "github.com/syukur91/ischool/storage/postgre"
	"github.com/syukur91/ischool/studentserver"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var (
	studentStorage domain.StudentStorage
)

func setupHandlers(e *echo.Echo, db *gorm.DB, log *log.Entry) {

	r := e.Group("/:tenant/api")

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello "+c.Param("tenant")+"! This is API version: "+os.Getenv("VERSION"))
	})

	studentStorage := postgreStorage.NewStudentPostgreStorage(db, log)

	studentserverHandler := &studentserver.Handler{
		StudentStorage: studentStorage,
		Log:            log,
	}

	studentserverHandler.SetRoutes(r)
}
