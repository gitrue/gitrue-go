package main

import (
	"github.com/astaxie/beego"
	_ "gitrue-go/models/orm"
	_ "gitrue-go/routers"
)

func main() {
	beego.Run()
}


