package main

import (
	"todo_list/conf"
	"todo_list/routes"
)

func main() {
	r := routes.NewRoute()
	_ = r.Run(conf.HttpPort)

}
