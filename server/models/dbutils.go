package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Resp struct {
	Items  interface{} `json:"items"`
	ErrMsg error       `json:"errMsg"`
}

const (
	UserName     = "example_name"
	Password     = "example_passwod"
	Addr         = "172.18.0.3"
	Port         = 3306
	MaxLifetime  = 10
	MaxOpenConns = 10
	MaxIdleConns = 10
)

type Product struct {
	//gorm.Model
	Id            int    `json:"id" gorm:"column:id"`
	Name          string `json:"name" gorm:"column:name"`
	Describe      string `json:"describe" gorm:"column:descr"`
	Price         int    `json:"price" gorm:"column:price"`
	Quantity      int    `json:"quantity" gorm:"column:quantity"`
	Series        int    `json:"series" gorm:"column:series"`
	Picture_path  string `json:"picture_path" gorm:"column:picture_path"`
	Created_time  string `json:"created_time" gorm:"column:created_time"`
	Modified_time string `json:"modified_time" gorm:"column:modified_time"`
	Deleted_time  string `json:"deleted_time" gorm:"column:deleted_time"`
}

type User struct {
	//gorm.Model
	Id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Location string `json:"location" gorm:"column:location"`
	PhoneNum string `json:"phone_num" gorm:"column:phone_num"`
}

//var connection *gorm.DB

func ConnectProductDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, "Product")

	// Connect to MySQL DB by gorm and set the config,
	// ignore the "s" at the end of the table name(SingularTable) and trun on the logger.
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("get db failed: ", err)
	}

	// Connect pool and set the pool
	db, err := conn.DB()
	if err != nil {
		log.Fatalln("get db failed: ", err)
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return conn
}

func ConnectUserDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, "User")

	// Connect to MySQL DB by gorm and set the config,
	// ignore the "s" at the end of the table name(SingularTable) and trun on the logger.
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("get db failed: ", err)
	}

	// Connect pool and set the pool
	db, err := conn.DB()
	if err != nil {
		log.Fatalln("get db failed: ", err)
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return conn
}

func GetProductALL(tName string) Resp {

	connection := ConnectProductDb()
	var res []Product
	err := connection.Table(tName).Find(&res).Error

	return Resp{res, err}
}

func GetProductByID(tName string, id string) Resp {

	connection := ConnectProductDb()

	var res []Product
	err := connection.Table(tName).Where("id = ?", id).Find(&res).Error

	return Resp{res, err}
}

func GetUserALL(tName string) Resp {

	connection := ConnectUserDb()
	var res []User
	err := connection.Table(tName).Find(&res).Error

	return Resp{res, err}
}

func GetUserByID(tName string, id string) Resp {

	connection := ConnectUserDb()

	var res []User
	err := connection.Table(tName).Where("id = ?", id).Find(&res).Error

	return Resp{res, err}
}

func GetUserByEmailPassword(info map[string]string) Resp {

	connection := ConnectUserDb()

	var res []User
	err := connection.Table("users").Where("email = ? and password = ?", info["email"], info["password"]).Find(&res).Error

	return Resp{res, err}
}
