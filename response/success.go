package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	StatusDesc string      `json:"status_desc"`
	Data 	 interface{} `json:"data"`
}

func SuccessHandler(c *fiber.Ctx, statusCode int, statusDesc string, data interface{}) error {
	logrus.Info(statusDesc)
    response := SuccessResponse{
        StatusCode: statusCode,
        StatusDesc: statusDesc,
        Data:       data,
    }
    return c.Status(statusCode).JSON(response)
}

type SuccessResponsePaginate struct {
	StatusCode int         `json:"status_code"`
	StatusDesc string      `json:"status_desc"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Size	   int         `json:"size"`
	TotalData  int         `json:"total_data"`
}

func SuccessHandlerPaginate(c *fiber.Ctx, statusCode int, statusDesc string, data interface{}, page int, size int, totalData int) error {
	logrus.Info(statusDesc)
    response := SuccessResponsePaginate{
        StatusCode: statusCode,
        StatusDesc: statusDesc,
        Data:       data,
        Page:       page,
        Size:       size,
        TotalData:  totalData,
    }
    return c.Status(statusCode).JSON(response)
}

