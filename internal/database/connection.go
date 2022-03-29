package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect(dsn string) *gorm.DB {
	//dsn := "host=localhost user=root password=root dbname=library port=3306" //sslmode=disable TimeZone=Asia/Shanghai

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "table_",
			SingularTable: true,
		},
		/*NowFunc: func() time.Time {
			return time.Now().UTC()
		},*/
	})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
	}

	DB = connection

	return DB
}
