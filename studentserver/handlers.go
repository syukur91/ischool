package studentserver

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/syukur91/ischool/domain"
)

type Handler struct {
	StudentStorage domain.StudentStorage
	Log            *log.Entry
}

// SetRoutes is
func (h *Handler) SetRoutes(r *echo.Group) {

	r.POST("/students", h.createStudent)

}

func (h *Handler) createStudent(c echo.Context) error {

	studentData := new(domain.Student)
	err := c.Bind(studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	students, err := h.StudentStorage.CreateStudent(studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, students)
}
