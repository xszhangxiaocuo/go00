/*
连接MySQL数据库,并修改记录
2022年11月30日21:19:15
*/

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 一个结构体对应一张数据表，一个结构体中的字段对应数据表中的字段,ID字段为默认主键，可用tab `gorm:"primaryKey"`修改主键
type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Age    int
}

func main() {
	//配置MySQL连接参数
	username := "xszxc"                                       //账号
	password := "Qaz6659644."                                 //密码
	host := "gz-cynosdbmysql-grp-hgzrzwmz.sql.tencentcdb.com" //数据库地址，可以是Ip或者域名
	port := 21105                                             //数据库端口
	Dbname := "demo"                                          //数据库名
	timeout := "10s"                                          //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //open方法会返回一个DB实例，此处并没有接收此参数
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	} else {
		print("successful")
	}

}
