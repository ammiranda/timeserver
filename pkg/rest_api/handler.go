package rest_api

import (
	"github.com/ammiranda/timeserver/pkg/rest_api/models/response"
	"github.com/ammiranda/timeserver/pkg/timestamp"
	"github.com/gin-gonic/gin"
)

var errorResponse = response.TimeResponseError{Error: "Invalid Date"}

func NewRouter(t timestamp.Service) *gin.Engine {
	r := gin.Default()

	r.GET("/api/", parseTime(t))
	r.GET("/api/:datetime", parseTime(t))

	return r
}

func parseTime(t timestamp.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		datetime := c.Param("datetime")
		unixTime, dateTime, err := t.ParseTime(datetime)
		if err != nil {
			c.JSON(400, errorResponse)
			return
		}

		resp := response.TimeResponse{Unix: unixTime, UTC: dateTime}

		c.JSON(200, resp)
	}
}
