package mysql

import (
	"github.com/fire-g/mark-go-db/db"
	"github.com/fire-g/mark-go-util/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//定义orm引擎
var (
	Engine     *xorm.Engine
	driverName = "mysql"
	Config     *db.DatabaseConfig
)

func InitMysql() *xorm.Engine {
	initMysql()
	return Engine
}

//创建orm引擎
func initMysql() {
	logger.Info.Println("初始化Mysql数据库...")
	var err error
	Engine, err = xorm.NewEngine(driverName,
		Config.Username+":"+Config.Password+"@tcp("+Config.Uri+")/"+Config.Database+"?charset=utf-8")
	if err != nil {
		logger.Fatal.Fatal("MySQL数据库连接失败:", err)
	}
	logger.Info.Println("数据库连接成功...")
}
