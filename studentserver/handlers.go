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
	r.GET("/student/:id", h.getStudent)
	r.GET("/student/:nis/nis", h.getStudentByNIS)
	r.PUT("/student/:id", h.updateStudent)
}

func (h *Handler) createStudent(c echo.Context) error {

	studentData := new(domain.Student)
	err := c.Bind(studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	studentData.SetCreated("syukur")
	students, err := h.StudentStorage.CreateStudent(studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, students)
}

func (h *Handler) getStudent(c echo.Context) error {

	id := c.Param("id")

	student, err := h.StudentStorage.GetStudent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, student)
}

func (h *Handler) getStudentByNIS(c echo.Context) error {

	nis := c.Param("nis")

	student, err := h.StudentStorage.GetStudentByNIS(nis)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, student)
}

func (h *Handler) updateStudent(c echo.Context) error {

	nis := c.Param("id")

	studentData := new(domain.Student)
	err := c.Bind(studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	studentData.SetUpdated("syukur")
	students, err := h.StudentStorage.UpdateStudent(nis, studentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, students)
}
