package db

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-log"
	"github.com/basic/config"
)

func initMysql() {

	var err error

	// 创建连接
	mysqlDB, err = gorm.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "hi_" + defaultTableName
	}

	mysqlDB.SingularTable(true)


	// 最大连接数
	//mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.DB().SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	//mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())
	mysqlDB.DB().SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}
}
