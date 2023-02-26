package helper

import (
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// add meta with type meta inside response and data with type interface because
// it's dynamic object
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//create response
func CreateResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	// if data exist return with data
	if data != nil {
		return Response{
			Meta: meta,
			Data: data,
		}
	}
	// if data not exist or nil
	return Response{
		Meta: meta,
	}
}

func SendResponse(c *gin.Context, message string, code int, status string, data interface{}) {
	c.JSON(
		code,
		CreateResponse(message,
			code,
			status,
			data))
}

func SendValidationErrorResponse(c *gin.Context, message string, code int, status string, err error) {

	errorText := FormatValidationError(err)

	errorResponse := gin.H{"errors": errorText}

	SendResponse(c, message, code, status, errorResponse)

	c.Error(err)
	c.Abort()
}

func SendErrorResponse(c *gin.Context, message string, code int, status string, err error, response interface{}) {
	var jsonResponse Response

	if err == nil && response == nil {
		// empty response error
		jsonResponse = CreateResponse(message, code, status, nil)
	} else if response == nil {
		// error with error type data
		errorResponse := gin.H{"errors": err.Error()}
		c.Error(err)
		jsonResponse = CreateResponse(message, code, status, errorResponse)
	} else {
		// custom error
		jsonResponse = CreateResponse(message, code, status, response)
	}

	c.Abort()
	c.JSON(code, jsonResponse)
}
