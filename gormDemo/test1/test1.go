/*
连接MySQL数据库
2022年11月30日16:54:49
*/

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//配置MySQL连接参数
	username := "xszxc"  //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 13306        //数据库端口
	Dbname := "demo"     //数据库名
	timeout := "10s"     //连接超时，10秒

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
