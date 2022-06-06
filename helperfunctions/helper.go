package helperfunctions

import (
	"RockPaper_Api/models"

	"github.com/gofiber/fiber/v2"
)

// this function returns a customized error and needs [context, errorcode, message]
func ReturnJSONError(c *fiber.Ctx, errorcode uint, message string) error {
	Response := models.ResponseObject{StatusCode: errorcode, Message: message} // package a JSON message for the user
	intcode := int(errorcode)                                                  // convert errorcode to int for c.Status function
	return c.Status(intcode).JSON(Response)                                    // return the response
}

// wrapper function that checks if error is nil and then executes return JSON error
func CheckError(c *fiber.Ctx, errorcode uint, err error, message string) {
	if err != nil {
		ReturnJSONError(c, errorcode, message)
	}
}

// this function returns a JSON response to the user needs [context, statuscode, message]
func ReturnJSONResponse(c *fiber.Ctx, statuscode uint, message string) error {
	PositiveResponse := models.ResponseObject{StatusCode: statuscode, Message: message}
	intstatus := int(statuscode)
	return c.Status(intstatus).JSON(PositiveResponse)
}
