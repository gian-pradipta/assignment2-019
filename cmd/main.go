package main

import (
	"rest_api_order/internal/router"
)

func main() {
	var router *router.Router = router.New()
	router.StartServer()
}
