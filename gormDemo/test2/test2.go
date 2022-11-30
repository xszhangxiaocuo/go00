/*
连接MySQL数据库,并创建一个表格和一条记录
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //open方法会返回一个DB实例，此处并没有接收此参数
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	} else {
		print("successful")
	}

	//创建表，自动迁移（将结构体与数据表进行对应）
	db.AutoMigrate(&UserInfo{})
	//初始化一个UserInfo的实例
	u1 := UserInfo{1, "新生张小搓", "男", 19}
	//将一个UserInfo的实例传入数据库，即在数据表中创建一行记录
	db.Create(&u1)

	var user UserInfo
	db.First(&user, 1)                   //根据整型主键查找
	fmt.Printf("%#v\n", user)            //%#v会按照go的语法格式输出
	db.First(&user, "name = ?", "新生张小搓") //根据字段name=“新生张小搓”进行查询
	fmt.Printf("%#v\n", user)
	//修改上文中获取的user记录的字段信息
	db.Model(&user).Updates(UserInfo{Name: "冰兔"})
	fmt.Printf("%#v\n", user)
	db.Model(&user).Updates(map[string]interface{}{"name": "冰兔兔兔兔", "gender": "女"})
	fmt.Printf("%#v\n", user)

	db.Delete(user)
}
