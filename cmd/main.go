package main

import ("github.com/AmirmahdiShahbazi/clean-web-api/data"
		"github.com/AmirmahdiShahbazi/clean-web-api/api"
)

func main() {
	api.InitRouters()
	data.ConnectToRedis()
	data.ConnectToPostgres()
}