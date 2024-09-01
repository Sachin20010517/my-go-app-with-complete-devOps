package api

import (
	"rest-api/dao"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

func Get_All_Students_API(c *fiber.Ctx) error {
	returnValue, err := dao.DB_GetAllStudent()

	if err != nil {
		return utils.SendErrorResponse(c,fiber.StatusBadRequest, err.Error())		
	}

	return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}
