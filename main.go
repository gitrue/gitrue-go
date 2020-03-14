package main

import (
	"github.com/astaxie/beego"
	_ "gitrue/models/orm"
	_ "gitrue/routers"
)

func main() {
	beego.Run()
}


