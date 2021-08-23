package main

import (
	"github.com/ammiranda/timeserver/pkg/rest_api"
	"github.com/ammiranda/timeserver/pkg/timestamp"
)

func main() {
	s := timestamp.NewService()

	router := rest_api.NewRouter(s)

	router.Run()
}
