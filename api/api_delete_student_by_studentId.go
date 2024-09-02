package api

import (
	"rest-api/dao"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteStudent_API (c *fiber.Ctx) error {
	studentId := c.Query("studentId")

	err := dao.DB_DeleteStudent(studentId)
	if err != nil{
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())

	}
	return utils.SendSuccessResponse(c)

}