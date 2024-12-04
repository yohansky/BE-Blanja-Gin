package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// dsn := "host=localhost user=yohanes password=root dbname=market port=5437 sslmode=disable TimeZone=Asia/Shanghai"
	for i := 0; i < 10; i++ {

		// formatnya = "postgres://<username>:<password>@<nama container di docker>:5432(port default)/<nama db>"
		DB, err = gorm.Open(postgres.Open("postgres://yohanes:root@db:5432/market"), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic(err)
	} else {
		log.Printf("Connected to DB")
	}
}
