package api

import (
	"rest-api/dao"
	"rest-api/dto"
	"rest-api/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rest-api/functions"
)

func CreateStudentApi(c *fiber.Ctx) error {

	inputObj := dto.Student{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	inputObj.StudentId = uuid.New().String()
	if err := functions.UniqueCheck(inputObj, "Students", []string{"StudentId"}); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if validationErr := validate.Struct(&inputObj); validationErr != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}
	err := dao.DB_CreateStudent(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}
