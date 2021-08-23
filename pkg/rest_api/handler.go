package rest_api

import (
	"github.com/ammiranda/timeserver/pkg/rest_api/models/response"
	"github.com/ammiranda/timeserver/pkg/timestamp"
	"github.com/gin-gonic/gin"
)

var errorResponse = response.TimeResponseError{Error: "Invalid Date"}

func NewRouter(t timestamp.Service) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/api/", cors, parseTime(t))
	r.GET("/api/:datetime", cors, parseTime(t))

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

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
}
