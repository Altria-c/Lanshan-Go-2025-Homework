package main

import (
	"lansan-learning/6/api"
	"lansan-learning/6/dao"
)

func main() {
	dao.LoadDatabase()
	api.InitRouter_gin()
}
