package studentserver

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/syukur91/ischool/student"
)

type Handler struct {
	StudentClient student.Student
	Log           *log.Entry
}

// SetRoutes is
func (h *Handler) SetRoutes(r *echo.Group) {

	r.GET("/students", h.getStudents)

}

func (h *Handler) getStudents(c echo.Context) error {

	students, err := h.StudentClient.GetStudents("3A")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, students)
}
