package api

import (
	"rest-api/dao"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

func Find_Student_By_StudentId_API(c *fiber.Ctx) error {
	studentId := c.Query("studentId")

	returnValue, err := dao.DB_FindStudentByStudentId(studentId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&returnValue)

}