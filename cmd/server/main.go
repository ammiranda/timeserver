package main

import (
	"github.com/ammiranda/timeserver/pkg/rest_api"
	"github.com/ammiranda/timeserver/pkg/timestamp"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	s := timestamp.NewService()

	router := rest_api.NewRouter(s)

	router.Use(cors.Default())

	router.Run()
}
