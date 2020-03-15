package orm

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"gitrue-go/models"
	_ "gitrue-go/models"
	"time"
)


func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/gitrue?charset=utf8", 30)
	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// register model
	orm.RegisterModel(new(models.User),new(models.UserAction),new(models.Card),new(models.Setting))
	//使用表名前缀
	//orm.RegisterModelWithPrefix("gitrue_", new(User))
	// create table
	orm.RunSyncdb("default", false, true)


	o := orm.NewOrm()
	user := &models.User{NickName: "slene",Sex:1,Empiric:0,Online:true,Private:false,UserAction:nil,Card:nil}

	// insert
	id, err := o.Insert(user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.NickName = "astaxie"
	num, err := o.Update(user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := &models.User{Id: user.Id}
	err = o.Read(u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	//num, err = o.Delete(u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

