package main

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api"
	"github.com/AmirmahdiShahbazi/clean-web-api/data"
)

func main() {
	api.InitRouters()
	data.ConnectToRedis()
	data.ConnectToPostgres()
}